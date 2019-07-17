package main

import "fmt"


type node struct {
   val int
   next *node
};

func main() {
  var val,number int
	//getting input from user to define size of the linked list
  fmt.Println("Enter a number for size of the list")
  fmt.Scan(&number)

	fmt.Println("Enter values for the linked list")
	//assigning temp value from user to head node
	fmt.Scan(&val)
  head := &node{val,nil}
  pin := head

  for index:=1; index<number;index++{
    fmt.Scan(&val)
    pin.next = &node{val,nil}
    pin = pin.next
  }


	fmt.Println("\n\n----Printing the forward linked list----\n")
	//loop starts with pin(pointer) set to the head node, loop prints the value of the node untill the next of the node is pointing to nil
	fmt.Print(" | ")
	for pin = head; pin != nil; pin=pin.next {
		fmt.Print(pin.val," | ")
		if (pin.next==nil){
			fmt.Print("end")
		}
	}


}
