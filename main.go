package main

import (
  "fmt"
  "math"
  "strings"
  "os"
)

var pho_map map[int]string

func convert(num int){
  var digit_len = int(math.Log10(float64(num))) + 1
  str_arr := make([]string, digit_len)
  for i:=0; i < digit_len; i++{
    str_arr[digit_len - 1 - i] = pho_map[num % 10]
    num /= 10
  }
  fmt.Println(strings.Join(str_arr[:], ""))
}

func main(){

  pho_map = make(map[int]string)

  pho_map[0] = "Zero"
  pho_map[1] = "One"
  pho_map[2] = "Two"
  pho_map[3] = "Three"
  pho_map[4] = "Four"
  pho_map[5] = "Five"
  pho_map[6] = "Six"
  pho_map[7] = "Seven"
  pho_map[8] = "Eight"
  pho_map[9] = "Nine"

  // for i:=0; i < len(os.Args); i++{
  //   convert(os.Args[i])
  // }

  // convert(45456756786785)
  // fmt.Println(5)
}
