package main

import "fmt"


type node struct {
   val int
   next *node
};

func main() {
    //reading an integer
    var number int
	var temp int

	var head node
	var pin *node
	
    fmt.Println("Enter a number for size of the list")
    fmt.Scan(&number)

	fmt.Scan(&temp)
	head.val = temp 
	
	pin = &head
	
	for index := 1; index < number; index++ {
		if(index!=number){
		pin.next = new(node)
		}else{
		pin.next = nil
		}
		pin = pin.next
		fmt.Scan(&temp)
		pin.val = temp 
//		fmt.Println(index)
	}

	
	for pin = &head; pin != nil; pin=pin.next {
		fmt.Println(pin.val)
	}

}
