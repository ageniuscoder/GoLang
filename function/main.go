package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func simple() {
	fmt.Println("This is a simple function.")
}

func add1(a int, b int) int {
	return a + b
}

func add2(a, b int) (result int) {
	result = a + b
	return
}
func main() {
	simple()
	add := add(5, 10)
	println("The sum is:", add)

	add1 := add1(10, 10)
	println("The sum is:", add1)

	add2 := add2(20, 30)
	println("The sum is:", add2)
}
