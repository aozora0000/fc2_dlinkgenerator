package main

import (
    "fmt"
    "net/http"
    "net/url"
    "os"
    "flag"
    "strconv"
    "./fc2"
    "regexp"
)

/* ----------------------- */
/* --- controller          */
/* ----------------------- */
func handler(w http.ResponseWriter, r *http.Request) {
    m, err := url.ParseQuery(r.URL.RawQuery)
    if err != nil {
        fmt.Println(err)
        return
    }
	if m["mid"] == nil {
        fmt.Fprintf(w,"mid parameter is required")
        return
    }
    var re = regexp.MustCompile("^[0-9]{8}[0-9a-zA-Z]{8}$")
    var mid = string(m["mid"][0])

    if !re.MatchString(mid) {
        fmt.Fprintf(w,"mid is not valid")
        return
    }
    detail := fc2.GetParams(mid)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Fprintf(w, detail)
}

/* ----------------------- */
/* --- main                */
/* ----------------------- */
func main() {
    // switch
    var (portNum int)
    flag.IntVar(&portNum, "port", 8080, "int flag")
    flag.IntVar(&portNum, "p", 8080, "int flag")
    flag.Parse()

    var port string
    port = ":"+strconv.Itoa(portNum)
    fmt.Println("listen port =", port)

    // route handler
    http.HandleFunc("/", handler)

    // do serve
    err := http.ListenAndServe(port, nil)

    // error abort
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
