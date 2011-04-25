package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "flag"
    "http"
    "log"
    "regexp"
)

func main() {
    flag.Parse()

    if flag.NArg() < 1 {
        fmt.Fprintf(os.Stderr, "Usage: %s [URL]\n", os.Args[0])
        os.Exit(1)
    }

    u := "http://kon.tl/create.php"
    d := map[string]string{"url": flag.Arg(0)}

    fmt.Println(" Long URL:", flag.Arg(0))

    r, err := http.PostForm(u, d)
    if err != nil {
        log.Fatalln(err)
    }
    defer r.Body.Close()

    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln(err)
    }

    s := string(b)
    x := regexp.MustCompile("<a href=(.*) target=_blank>")
    m := x.FindStringSubmatch(s)[1]

    fmt.Println("Short URL:", m)
}
