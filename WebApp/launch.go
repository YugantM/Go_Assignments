package main

import ("fmt"
        "net/http")

func index_page(response http.ResponseWriter,request *http.Request){
  fmt.Fprintf(response, "Hello there!")
}

func main(){
  http.HandleFunc("/",index_page)
  http.ListenAndServe(":8000",nil)
}
