package main

import ("fmt"
        "log"
        "net/http")

// the handler function for "/" route/home
func index_page(response http.ResponseWriter,request *http.Request){
  fmt.Fprintf(response, "Hello there!")
}

func main(){

  //it routes the given handler function
  http.HandleFunc("/",index_page)

  //starting the server on desired port, second argument is for handler which is nil
  log.Fatal(http.ListenAndServe(":8000",nil))
}
