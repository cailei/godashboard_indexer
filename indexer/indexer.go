/*
gopm indexer of 'http://godashboard.appspot.com/' (Go Package Manager)
Copyright (c) 2012 cailei (dancercl@gmail.com)

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
    "flag"
    "fmt"
    "github.com/hailiang/gosocks"
    "io/ioutil"
    "log"
    "net/http"
)

var remote_db_host string = "http://localhost:8080"

func main() {
    log.SetFlags(log.Lshortfile)

    // parse flags
    var proxy_addr string
    var help bool
    flag.StringVar(&proxy_addr, "socks5-proxy", "", "use proxy to access the server")
    flag.StringVar(&help, "help", false, "show help")
    flag.StringVar(&help, "h", false, "show help")
    flag.Usage = print_usage
    flag.Parse()

    client := get_http_client(proxy_addr)
    page := get_dashboard_page(client)

    fmt.Print(page)
}

func get_dashboard_page(client *http.Client) string {
    request := "http://godashboard.appspot.com/"

    // request the page
    response, err := client.Get(request)
    if err != nil {
        log.Fatalln(err)
    }

    // check response
    if response.StatusCode != 200 {
        page, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatalln(err)
        }

        if len(page) > 0 {
            log.Fatalln(string(page))
        }

        log.Fatalln(response.Status)
    }

    // get page content
    page, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }

    return string(page)
}

func get_http_client(proxy_addr string) *http.Client {
    client := http.DefaultClient

    // check if using a proxy
    if proxy_addr != "" {
        proxy := socks.DialSocksProxy(socks.SOCKS5, proxy_addr)
        transport := &http.Transport{Dial: proxy}
        client = &http.Client{Transport: transport}
    }

    return client

}

func print_usage() {
    fmt.Print(`
indexer:

options:
    -socks5-proxy   use a socks5 proxy to access the net
    -h, -help       show help info

`)
}
