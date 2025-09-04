package main //chanenels is used to communicate between goroutines on the same program
import "fmt"

//sending data to channel
func sum(channel chan int, a int, b int) {
	channel <- a + b // Send the sum of a and b to the channel
}

// for recieving data from a channel
//
//	func processChannel(channel chan int) {
//		for value := range channel {
//			fmt.Println("Received:", value) // Print the received value from the channel
//			time.Sleep(1 * time.Second)     // Simulate some processing time
//		}
//	}
func main() {
	channel := make(chan int) // Create a channel of type int
	go sum(channel, 5, 10)    // Start a goroutine that will send the sum to the channel
	result := <-channel

	fmt.Println("Sum:", result) // Print the result received from the channel

	// channel := make(chan int)  // Create a channel of type int
	// go processChannel(channel) // Start a goroutine that will process the channel

	// for {
	// 	channel <- rand.Intn(100) // Send a random integer to the channel
	// }
	// channel := make(chan int)
	// channel <- 42    // This will cause a deadlock because no goroutine is reading from the channel
	// msg := <-channel // This line will never be reached
	// fmt.Println(msg) // This line will never be executed

}
