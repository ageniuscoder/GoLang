package main

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 42
	}()

	go func() {
		chan2 <- "Hello, World!"
	}()

	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			println("Received from chan1:", num)
		case str := <-chan2:
			println("Received from chan2:", str)
		}
	}

}
