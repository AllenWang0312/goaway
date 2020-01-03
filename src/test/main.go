package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	message := make(chan string)
	wg.Add(1)
	go func() {
		wg.Done()
		message <- "hello goroutine"
		//fmt.Println("hello goroutine")
	}()
	fmt.Println(<-message)
	fmt.Println("hello world")
	wg.Wait()
}
