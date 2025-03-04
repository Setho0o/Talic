package main

import "fmt"

func main() {
  path := "audio/The_Pretender.wav" 
  m := GetMetaData(path)
  fmt.Println(m)
}
