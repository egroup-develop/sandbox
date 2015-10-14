package main

import(
  "fmt"
  //"bufio"
  //"os"
)

//ユーザーにやらせる部分ここ. これから実装する
func merge(a, b []int)[]int{
  tmp := make([]int, len(a)+len(b))
  i, j := 0, 0

  for i < len(a) && j < len(b){
    if a[i] <= b[j]{
      tmp[i+j] = a[i]
      i++
    }else{
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
  a := []int{10, 9, 8, 4, 6, 5, 6, 7, 3, 9, 2, 1}
  
  //標準入力
  /*scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan(){
    fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil{
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
  }*/
  //マージソート
  fmt.Println(mergeSort(a))
}
