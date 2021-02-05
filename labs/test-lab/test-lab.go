package main

import (
	"fmt"
	"os"
)

func main() {

  if len(os.Args)==1{
    fmt.Println("Error")
  }
  if len(os.Args)>1{
    fmt.Print("Hello")
    for _,word := range os.Args[1:]{
      fmt.Print(" ",word)
    }
    fmt.Println(", Welcome to the Jungle")
  }
  
}