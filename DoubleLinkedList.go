package main

import "fmt"


type node struct {
   val int
   next,previous *node
};

func main() {
    //variables
  var number int
	var temp int

	//nodes (structures)
	var head,tail node
	var pin *node

	//getting input from user to define size of the linked list
  fmt.Println("Enter a number for size of the list")
  fmt.Scan(&number)

	fmt.Println("Enter values for the linked list")
	//assigning temp value from user to head node
	fmt.Scan(&temp)
	head.val = temp
  head.previous = nil

	//pointer pointing to the head node
	pin = &head
  fmt.Println("---printing head::---\n)
	//loop for adding nodes
	for index := 1; index < number; index++ {

		//checking if the loop is on the last number or not
		if(index!=number){

    pin.previous = pin
    tail = *pin
    pin.next = new(node)
    // if not on last number, allocating new node to the next of the current one

  	}else{
		pin.next = nil	      // if yes allocate nil to the next of the last node
		}

		//getting values from user and adding it to the current node
		pin = pin.next
		fmt.Scan(&temp)
		pin.val = temp
    fmt.Println("prev:",pin.previous,"\nval",pin.val,"\nnext:",pin.next,"\ntail\n\n:",&tail,"\n\n")
	}

/*  tail = pin.previous

	fmt.Println("Printing the forward linked list")
	//loop starts with pin(pointer) set to the head node, loop prints the value of the node untill the next of the node is pointing to nil
	fmt.Print(" | ")
	for pin = &head; pin != nil; pin=pin.next {
		fmt.Print(pin.val," | ")
		if (pin.next==nil){
			fmt.Print("end")
		}
	}
*/

  //fmt.Println("\n\nPrinting the backward linked list")

//  fmt.Println("garbage\n\n",tail.val,tail.previous,tail.next)
	//loop starts with pin(pointer) set to the head node, loop prints the value of the node untill the next of the node is pointing to nil
	//fmt.Print(" | ")
/*	for pin = &tail; pin != nil; pin=pin.previous {
		fmt.Print(pin.val," | ")
		if (pin.previous==nil){
			fmt.Print("start")
		}
	}*/


}
