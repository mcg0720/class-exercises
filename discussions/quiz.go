package main

import (
	"fmt"
	"sync"
)

// func worker(ch chan int) {
// 	i := <-ch
// 	ch <- i * 2
// }

var mu sync.Mutex

func worker(ch chan int) {
	//mu.Lock()
	//defer mu.Unlock()
	i := <-ch
	ch <- i * 2
}

func main() {
	// ch := make(chan int)
	// go worker(ch)
	// ch <- 3
	// fmt.Println(<-ch)

	ch := make(chan int)
	mu.Lock()
	ch <- 1
	mu.Unlock()
	go worker(ch)
	fmt.Println(<-ch)
}
