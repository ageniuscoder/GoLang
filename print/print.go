package main

import "fmt"

func main() {
	var name string = "Mangal"                                              // var declaration using datatype making it explicit
	age := 30                                                               // short variable declaration, only works inside functions
	height := 5.97463                                                       // short variable declaration for float
	fmt.Println("My name is", name, "age is", age, "and height is", height) //using this i can,t do fomatting of data according to my need it always provide new lines and spaces intelligently

	//for formatted output i can use Printf
	fmt.Printf("My name is %s\n, age is %d\n, and height is %.2f\n", name, age, height) // using Printf for formatted output
}
