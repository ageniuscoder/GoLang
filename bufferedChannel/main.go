package main

import (
	"fmt"
	"time"
)

// <-chan recieve only                              chan <- send only
func sendEmails(email <-chan string, done chan<- bool) { //for safety , i can make email channel reacieve only channel and done channel send only channel
	defer func() { done <- true }()
	for e := range email {
		fmt.Println("Sending email to:", e)
		time.Sleep(1 * time.Second) // Simulate time taken to send an email
	}
}

func main() {
	email := make(chan string, 50)
	done := make(chan bool)
	go sendEmails(email, done)
	for i := 1; i <= 6; i++ {
		email <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("All emails sent, closing channel.")
	close(email)
	<-done
	fmt.Println("All emails processed and delivered, exiting program.")
}
