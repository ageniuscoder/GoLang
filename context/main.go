package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	genrator := func(dataItem string, stream chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- dataItem:
			}
		}
	}

	infinitApples := make(chan interface{})
	wg.Add(1)
	go genrator("apple", infinitApples)

	infinitBanana := make(chan interface{})
	go genrator("banana", infinitBanana)

	infinitMango := make(chan interface{})
	go genrator("mango", infinitMango)

	wg.Add(1)
	go func1(ctx, &wg, infinitApples)

	func2 := genricFunc
	func3 := genricFunc

	wg.Add(1)
	go func2(ctx, &wg, infinitBanana)

	wg.Add(1)
	go func3(ctx, &wg, infinitMango)

	wg.Wait()

}

func func1(ctx context.Context, parentWg *sync.WaitGroup, stream <-chan interface{}) {
	defer parentWg.Done()

	var wg sync.WaitGroup

	dowork := func(ctx context.Context) {
		wg.Done()

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

	newCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go dowork(newCtx)
	}

	wg.Wait()
}

func genricFunc(ctx context.Context, wg *sync.WaitGroup, stream <-chan interface{}) {
	wg.Done()

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
