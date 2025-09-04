package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}
func main() {
	fmt.Println("Error handling example in Go")
	result, _ := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Result:", result)
	// }

	fmt.Println("Result:", result)
}
