package main //use chaneel as waitGroup synchronization
import (
	"fmt"
	"time"
)

func syncChannel(done chan bool) {
	defer func() { done <- true }() // signal completion
	fmt.Println("syncChannel: started")
	time.Sleep(2 * time.Second) // simulate work
}
func main() {
	done := make(chan bool)

	go syncChannel(done) // start the goroutine

	<-done // wait for the goroutine to finish  it blocks until the channel receives a value
	fmt.Println("syncChannel: finished")
}
