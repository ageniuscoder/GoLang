package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Enter Your Name:")
	// var name string
	// fmt.Scan(&name) //read input till whitespace
	// fmt.Println("Hello,", name, "! Welcome to the program.")
	// var line string
	// fmt.Println("Enter a line of text:")
	// fmt.Scanln(&line) //read input till newline
	// fmt.Println("You entered:", line)
	fmt.Println("Enter a sentence:")
	reader := bufio.NewReader(os.Stdin) //create a reader object

	line, _ := reader.ReadString('\n') //read input till newline

	fmt.Println("You entered:", line)
}
