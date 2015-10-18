/**
 * このファイルは, issue#5の課題を達成するためのもの
 */

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
/*** Duration()使う時にコメント外す
  "time"
****/
  "encoding/json"
  "bytes"
)

const (
  //リクエストを投げるURL
  REQ_URL = "http://weather.livedoor.com/forecast/webservice/json"
)

//リクエストのレスポンス格納用
var responseGet, responsePost string = "", ""

/*** JSONファイルと同じ構造体定義ここから ***/
type image struct{
  Title string
  Url string
}

type forecasts struct{
  Date string
  Image image
}

type data struct{
  Forecasts []forecasts
}
/*** JSONファイルと同じ構造体定義ここまで ***/

/*** 時間制限ありGETを使うときコメントを外す
func get(client *http.Client, values url.Values) {
  req, err := http.NewRequest("GET", REQ_URL + "/v1", nil)

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

  execute(resp, "GET")
}
***/

func get(values url.Values){
  resp, err := http.Get(REQ_URL + "/v1?" + values.Encode())

  if err != nil{
    fmt.Println(err)
    return
  }
  defer resp.Body.Close()

  execute(resp, "GET")
}

func post(values url.Values){
  resp, err := http.PostForm("http://httpbin.org/post", values)

  if err != nil{
    fmt.Println(err)
    return
  }
  defer resp.Body.Close()

  execute(resp, "POST")
}

func execute(resp *http.Response, kinds string) {
  //response bodyを文字列で取得
  b, err := ioutil.ReadAll(resp.Body)

  if err == nil {
    if kinds == "GET"{
      responseGet = string(b)
    }else if kinds == "POST"{
      responsePost = string(b)
    }
  }
}

func main() {
  values := url.Values{}

  //リクエストパラメータ用のkey-valueを追加
  values.Add("city", "200010")

  fmt.Println("パラメータ:\n", values.Encode())

/***
  //時間制限ありGET
  client := &http.Client{Timeout: time.Duration(10 * time.Second)}
  get(client, values)
****/

  //普通のGET
  get(values)
  //普通のPOST
  post(values)

  fmt.Println("GETリクエスト:\n", responseGet)
  fmt.Println("POSTリクエスト:\n", responsePost)

  //受け取るJSONファイルの解析
  dec := json.NewDecoder(bytes.NewBufferString(responseGet))
  var d data
  dec.Decode(&d)

  fmt.Println("【date: 日付, title: 天気, url: URL】")
  fmt.Printf("%+v\n", d)
}
