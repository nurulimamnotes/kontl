package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "flag"
    "http"
    "regexp"
)

func main() {
    flag.Parse()

    if flag.NArg() < 1 {
        fmt.Fprintln(os.Stderr, "Usage: kontl [URL]")
        os.Exit(-1)
    }

    u := "http://kon.tl/create.php"
    d := map[string]string{"url": flag.Arg(0)}

    fmt.Println("Long URL:", flag.Arg(0))

    r, err := http.PostForm(u, d)
    if err != nil {
        fmt.Println("POST:", err)
        os.Exit(-1)
    }
    defer r.Body.Close()

    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println("POST:", err)
        os.Exit(-1)
    }

    s := string(b)
    x := regexp.MustCompile("<a href=(.*) target=_blank")
    m := x.FindStringSubmatch(s)[1]

    fmt.Println("Short URL:", m)
}
