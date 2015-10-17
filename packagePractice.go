/**
 * このファイルは, issue#5の課題を達成するためのもの
 */


//POST リクエスト

package main

import (
    "net/http"
    "net/url"
    "io/ioutil"
)

func main() {

    resp, _ := http.PostForm(
        "http://weather.livedoor.com/forecast/webservice/json/v1?city=200010",
        url.Values{"foo": {"bar"}, },
    )

    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    println(string(body))
}
