/**
 * Charソート(マージ)完成版!!!
 * 結果は好みの昇順で表示しまっっっ!!!
 */

package main

import(
  "fmt"
  "bufio"
  "os"
)

//ユーザーに好きな数字を選択させる
func merge(a, b []int)[]int{
  tmp := make([]int, len(a)+len(b))
  i, j := 0, 0
  eval := 0

  for i < len(a) && j < len(b){
    //ここで入力要求
    fmt.Println(a[i], " と ", b[j], " どっちの数字が好き?")
    fmt.Println(a[i], " なら 0 を, ", b[j], " なら 1 を入力してね. 例外処理していないので0と1以外は絶対に入力しないでね.")
    //標準入力
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan(){
      switch scanner.Text(){
        case "0":
          eval = 0
        case "1":
          eval = 1
        default:
      }
  
      //fmt.Println(eval)
    }
    //エラーハンドリング
    if err := scanner.Err(); err != nil{
      fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

    if eval == 1{
      tmp[i+j] = a[i]
      i++
    }else if eval == 0{
      tmp[i+j] = b[j]
      j++
    }
  }

  for i < len(a){
    tmp[i+j] = a[i]
    i++
  }

  for j < len(b){
    tmp[i+j] = b[j]
    j++
  } 

  return tmp
}

func mergeSort(items []int)[]int{
  if len(items) > 1{
    return merge(mergeSort(items[:len(items)/2]), mergeSort(items[len(items)/2:]))
  }

  return items
}

func main(){
  a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

  //マージソート
  fmt.Println(mergeSort(a))
}
