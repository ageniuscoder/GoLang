package main

import "fmt"

// func PrintSlices[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func PrintSlices[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
func main() {
	num := []int{1, 2, 3}
	PrintSlices(num)

	strings := []string{"hello", "world"}
	PrintSlices(strings)

	bools := []bool{true, false, true}
	PrintSlices(bools)
}
