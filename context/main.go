package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// Generic generator function
	genrator := func(dataItem string, stream chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- dataItem:
			}
		}
	}

	// Create channels
	infinitApples := make(chan interface{})
	infinitBanana := make(chan interface{})
	infinitMango := make(chan interface{})

	// Start generators
	go genrator("apple", infinitApples)
	go genrator("banana", infinitBanana)
	go genrator("mango", infinitMango)

	// Consumers
	wg.Add(1)
	go func1(ctx, &wg, infinitApples) // apple consumer will stop after 5ms

	wg.Add(1)
	go genricFunc(ctx, &wg, infinitBanana) // keep running till 10ms

	wg.Add(1)
	go genricFunc(ctx, &wg, infinitMango) // keep running till 10ms

	wg.Wait()
}

func func1(ctx context.Context, parentWg *sync.WaitGroup, stream <-chan interface{}) {
	defer parentWg.Done()

	var wg sync.WaitGroup

	// Create a 5ms timeout context for apples only
	newCtx, cancel := context.WithTimeout(ctx, 5*time.Millisecond)
	defer cancel()

	dowork := func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-stream:
				if !ok {
					fmt.Println("apple channel closed")
					return
				}
				fmt.Println(d)
			}
		}
	}

	// Start 3 apple consumers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go dowork(newCtx)
	}

	wg.Wait() // wait until all apple consumers finish
}

func genricFunc(ctx context.Context, wg *sync.WaitGroup, stream <-chan interface{}) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-stream:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(d)
		}
	}
}
