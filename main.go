package main

import (
  "fmt"
  "strings"
  "os"
)

var pho_map map[string]string

func convert(str_num string){
  str_arr := make([]string, len(str_num))
  for i:=0; i < len(str_num); i++{
    str_arr[i] = pho_map[string(str_num[i])]
  }

  fmt.Print(strings.Join(str_arr[:], ""))
}

func main(){

  pho_map = make(map[string]string)

  pho_map["0"] = "Zero"
  pho_map["1"] = "One"
  pho_map["2"] = "Two"
  pho_map["3"] = "Three"
  pho_map["4"] = "Four"
  pho_map["5"] = "Five"
  pho_map["6"] = "Six"
  pho_map["7"] = "Seven"
  pho_map["8"] = "Eight"
  pho_map["9"] = "Nine"

  argsWithoutProg := os.Args[1:]
  for i:=0; i < len(argsWithoutProg); i++{
    convert(argsWithoutProg[i])
    if i != len(argsWithoutProg) - 1{
      fmt.Print(",")
    }
  }

}
