package fc2

import(
    "crypto/md5"
    "fmt"
    "net/http"
    "net/url"
    "strconv"
    "io"
    "io/ioutil"
    "encoding/json"
    "runtime"
)

type Detail struct {
    Iserror         bool
    Path            string
    Title           string
    Isadult         bool
    Payment         bool
    Length          int
}

var error_json = "{Iserror: true}"

func init() {
    cpus := runtime.NumCPU()
    runtime.GOMAXPROCS(cpus)
}

func mimi(video_id string) string {
    h := md5.New()
    io.WriteString(h, video_id + "_gGddgPfeaf_gzyr")
    return fmt.Sprintf("%x", h.Sum(nil))
}

func GetParams(video_id string) string {
    url := fmt.Sprintf("http://video.fc2.com/ginfo.php?mimi=%s&upid=%s",mimi(video_id),video_id)
    return get_request(url)
}

func get_request(url string) string {
    req, _ := http.NewRequest("GET", url, nil)
    client := new(http.Client)
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    byteArray, _ := ioutil.ReadAll(resp.Body)
    return decode_params(string(byteArray))
}

func decode_params(params string) string {
    u, _ := url.Parse("http://localhost/?" + params)
    m, err := url.ParseQuery(u.RawQuery)
    if (err != nil) {
        fmt.Println(err)
        return error_json
    }
    if m["err_code"][0] == "404" {
        return error_json
    }

    switch {
        case !validate(m, "err_code"):
        case !validate(m, "filepath"):
        case !validate(m, "mid"):
        case !validate(m, "title"):
        case !validate(m, "isadult"):
        case !validate(m, "payment"):
        case !validate(m, "sec"):
            return error_json
    }

    var err_code,_   = strconv.ParseBool(m["err_code"][0])
    var filepath     = m["filepath"][0]
    var mid          = m["mid"][0]
    var title        = m["title"][0]
    var isadult,_    = strconv.ParseBool(m["isadult"][0])
    var payment,_    = strconv.ParseBool(m["payment"][0])
    var sec,_        = strconv.Atoi(m["sec"][0])
    var downloadpath = filepath + "?mid=" + mid

    detail := &Detail{Iserror: err_code, Path: downloadpath, Title: title, Isadult: isadult, Payment: payment, Length: sec}
    b, err := json.Marshal(detail)
    if err != nil {
        fmt.Println(err)
        return "error"
    }
    return string(b)
}

func validate(m url.Values, key string) bool {
    _, ok := m[key]
    return ok
}
