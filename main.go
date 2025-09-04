package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hey i am mangal")
	myutil.PrintMessage("Hello from myutil package!")

	//learning variables
	var name string = "Mangal" //var decleration using datatype making it explicit
	fmt.Println("My name is", name)

	var name1 = "Mangal Patel" //var decleration using datatype inference
	fmt.Println("My name is", name1)

	name2 := "Mangal with go" //short variable declaration, only works inside functions cant use in other packages
	fmt.Println("My name is", name2)

	var PublicName = "Mangal Patel" //public variable, accessible outside the package
	fmt.Println("My public name is", PublicName)

	var privateName = "Mangal Patel"               //private variable, not accessible outside the package
	fmt.Println("My private name is", privateName) //same for functions, if the first letter is lowercase, it is private to the package if it is uppercase, it is public

	const pi = 3.14 //constant variable, cannot be changed
	fmt.Println("Value of pi is", pi)
}
