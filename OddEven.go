package main

import "fmt"

func main() {
    //reading an integer
    var number int
    fmt.Println("Enter a number")
    fmt.Scan(&number)

	if(number%2==0){
		fmt.Println("Even")
	}else if(number%2!=0){
		fmt.Println("Odd")
	}else{
		fmt.Println("What you are even typing !!")
	}

}
