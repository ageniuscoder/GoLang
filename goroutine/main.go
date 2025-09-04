package main

import (
	"fmt"
	"time"
)

func f1() {
	fmt.Println("hey it,s starting")
	time.Sleep(2 * time.Second)
	fmt.Println("hey it,s ending")
}

func f2() {
	fmt.Println("hey it,s testing")
	time.Sleep(1 * time.Second)
	fmt.Println("hey it,s done")
}

func main() {
	go f1()
	go f2()
	time.Sleep(1 * time.Second)
}
