package main

import "fmt"


type node struct {
   val int
   next *node
};

func main() {
    //variables
    var number int
	var temp int

	//nodes (structures)  
	var head node
	var pin *node
	
	//getting input from user to define size of the linked list
    fmt.Println("Enter a number for size of the list")
    fmt.Scan(&number)

	fmt.Println("Enter values for the linked list")
	//assigning temp value from user to head node
	fmt.Scan(&temp)
	head.val = temp 
	
	//pointer pointing to the head node
	pin = &head
	
	//loop for adding nodes 
	for index := 1; index < number; index++ {
		
		//checking if the loop is on the last number or not
		if(index!=number){
		pin.next = new(node)  // if not on last number, allocating new node to the next of the current one
		}else{
		pin.next = nil	      // if yes allocate nil to the next of the last node
		}
		
		//getting values from user and adding it to the current node
		pin = pin.next
		fmt.Scan(&temp)
		pin.val = temp 

	}

	fmt.Println("Printing the linked list")
	//loop starts with pin(pointer) set to the head node, loop prints the value of the node untill the next of the node is pointing to nil
	fmt.Print(" | ")
	for pin = &head; pin != nil; pin=pin.next {
		fmt.Print(pin.val," | ")
		if (pin.next==nil){
			fmt.Print("end")
		}
	}

}
