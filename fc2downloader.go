package main

import(
    "crypto/md5"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
    "path"
    "os"
    "bufio"
    "sync"
    "runtime"
)


type Detail struct {
    err_code        int
    filepath        string
    mid             string
    title           string
    isadult         int
    payment         int
    sec             int
}

func main() {
    //引数のパース
    flag.Parse()
    var arg = flag.Args()

    //ファイルパース
    lines := getfile(arg[0])

    var wg sync.WaitGroup
    cpus := runtime.NumCPU()
    //runtime.GOMAXPROCS(cpus)

    semaphore := make(chan int, cpus)
    for i, value := range lines {
        wg.Add(1)
        go func(i int, value string) {
            defer wg.Done()
            semaphore <- 1
            detail := params_uri(value)
            fmt.Println(detail.filepath + "?mid=" + detail.mid)
            <-semaphore
        }(i,value)
    }
    wg.Wait()
}

func mimi(video_id string) string {
    h := md5.New()
    io.WriteString(h, video_id + "_gGddgPfeaf_gzyr")
    return fmt.Sprintf("%x", h.Sum(nil))
}

func params_uri(video_id string) *Detail {
    url := fmt.Sprintf("http://video.fc2.com/ginfo.php?mimi=%s&upid=%s",mimi(video_id),video_id)
    return get_request(url)
}

func get_request(url string) *Detail {
    req, _ := http.NewRequest("GET", url, nil)
    client := new(http.Client)
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    byteArray, _ := ioutil.ReadAll(resp.Body)
    return decode_params(string(byteArray))
}

func decode_params(params string) *Detail {
    u, _ := url.Parse("http://localhost/?" + params)
    m, _ := url.ParseQuery(u.RawQuery)

    detail := new(Detail)
    detail.err_code,_  = strconv.Atoi(m["err_code"][0])
    detail.filepath    = m["filepath"][0]
    detail.mid         = m["mid"][0]
    detail.title       = m["title"][0]
    detail.isadult,_   = strconv.Atoi(m["isadult"][0])
    detail.payment,_   = strconv.Atoi(m["payment"][0])
    detail.sec,_       = strconv.Atoi(m["sec"][0])
    return detail
}

func getfile(filePath string) []string {
    f, err := os.Open(filePath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
        os.Exit(1)
    }

    // 関数return時に閉じる
    defer f.Close()

    // Scannerで読み込む
    lines := make([]string, 0, 1000)
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if serr := scanner.Err(); serr != nil {
        fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filePath, err)
    }
    return lines
}


func download(detail *Detail) {
    var url string = detail.filepath + "?mid=" + detail.mid

    response, err := http.Get(url)
    if err != nil {
        println(err)
    }
    defer response.Body.Close()

    //ステータスの表示
    fmt.Println("status:", response.Status)

    //ファイルを開く
    file, err := os.OpenFile("./" + path.Base(detail.filepath), os.O_CREATE|os.O_WRONLY, 0777)

    //エラーチェック
    if err != nil {
        panic(err)
    }

    //ファイルを閉じる関数の遅延実行指定
    defer file.Close()

    //レスポンスのボディから読み込みつつファイルに書き出す。
    io.Copy(file, response.Body)
}
