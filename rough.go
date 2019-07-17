package main

import "fmt"

type node struct {
   previous *node
   val int
   next *node
};

func main(){

  head := &node{nil,789,nil}
  head.previous = nil
  x := head
  temp := head


  for index:=1; index<5;index++{
    x.next = &node{nil,index,nil}
    temp = x
    x = x.next
    x.previous = temp
  }
    tail := x




  fmt.Println("\n\n----Forward-----")
  for pin := head; pin != nil; pin=pin.next {
  		fmt.Print("\n",pin," | ")
  		if (pin.next==nil){
  			fmt.Print("end")
  		}
    }



  fmt.Println("\n\n----Reverse-----")

  for pin := tail; pin != nil; pin=pin.previous {
        fmt.Print("\n",pin," | ")
        if (pin.previous==nil){
          fmt.Print("start")
        }
      }

}
