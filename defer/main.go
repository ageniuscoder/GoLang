package main

import "fmt"

func main() {
	fmt.Println("Starting the defer example...")
	defer fmt.Println("This will be printed last, after the main function completes.")
	defer fmt.Println("This will be printed second, right before the main function exits.")
	fmt.Println("This will be printed first, before the defer statement.")
}

//when there is multiple defer statements, they are executed in LIFO (Last In, First Out) order.
