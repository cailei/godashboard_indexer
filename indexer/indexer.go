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
    flag.StringVar(&proxy_addr, "proxy", "", "use proxy to access the server")
    flag.Usage = print_usage
    flag.Parse()

    request := "http://godashboard.appspot.com/"

    client := http.DefaultClient

    // check if using a proxy
    if proxy_addr != "" {
        proxy := socks.DialSocksProxy(socks.SOCKS5, proxy_addr)
        transport := &http.Transport{Dial: proxy}
        client = &http.Client{Transport: transport}
    }

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

    fmt.Print(string(page))
}

func print_usage() {
    fmt.Print(`
gopm create <package>:
    this wil create a <package.json> file containing information for your
    package, you should modify this file to fill in the fields manually, then
    run 'gopm publish <package.json>' to upload the information to the index
    server.

options:
    -f, -force      force overwrite existing file
    -h, -help       show help info

`)
}
