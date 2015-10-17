/**
 * このファイルは, issue#5の課題を達成するためのもの
 */

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "time"
)

const (
  //リクエストを投げるURL
  REQ_URL = "http://weather.livedoor.com/forecast/webservice/json"
)

//GETリクエストのレスポンスを格納用
var response string = ""

func main() {
  //url.Valuesオブジェクト生成
  values := url.Values{}

  //パラメータ用のkey-valueを追加
  values.Add("city", "200010")
  fmt.Println(values.Encode())

  client := &http.Client{Timeout: time.Duration(10 * time.Second)}
  get(client, values)

  fmt.Println(response)
}

func get(client *http.Client, values url.Values) {
  req, err := http.NewRequest("GET", REQ_URL+"/v1", nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  req.URL.RawQuery = values.Encode()

  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer resp.Body.Close()

  execute(resp)
}

func execute(resp *http.Response) {
  //response bodyを文字列で取得するサンプル
  //ioutil.ReadAllを使う
  b, err := ioutil.ReadAll(resp.Body)
  if err == nil {
    //fmt.Println(string(b))
    response = string(b)
  }
}
