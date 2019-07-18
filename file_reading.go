package main

import("fmt"
       "io/ioutil"
       "strings"
       )

func main(){

  data, err := ioutil.ReadFile("text.txt")
      if err != nil {
          fmt.Println("File reading error", err)
          return
      }
      //fmt.Println("Contents of file:", string(data))

      rows:= strings.Split(string(data),"\n")
      fmt.Println("\n\nContents of file:", rows[1])
}
