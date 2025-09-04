package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu    sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	p.mu.Lock()
	defer p.mu.Unlock() //best practice to put jsut after lock
	p.views += 1

}

func main() {
	var wg sync.WaitGroup
	mypost := Post{
		views: 0,
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go mypost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(mypost.views)
}
