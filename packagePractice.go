/**
 * このファイルは, issue#5の課題を達成するためのもの
 */

package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
/*Duration()使う時にコメント外す
  "time"
*/
)

const (
  //リクエストを投げるURL
  REQ_URL = "http://weather.livedoor.com/forecast/webservice/json"
)

//リクエストのレスポンス格納用
var responseGet, responsePost string = "", ""

/*時間制限ありGETを使うときコメントを外す
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
*/


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
    //fmt.Println(string(b))
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
  fmt.Println(values.Encode())

/*
  //時間制限ありGET
  client := &http.Client{Timeout: time.Duration(10 * time.Second)}
  get(client, values)
*/

  //普通のGET
  get(values)

  //普通のPOST
  post(values)

  fmt.Println(responseGet)
  fmt.Println(responsePost)
}
