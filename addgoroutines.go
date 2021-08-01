package main

import (
	"fmt"
	"sync"
)

const xthreads = 10

func doSomething(a int) {
	fmt.Println("My channel is", a)
	return
}

func main() {
	var ch = make(chan int, 50)
	var wg sync.WaitGroup

	wg.Add(xthreads)
	for i := 0; i < xthreads; i++ {
		go func() {
			for {
				a, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				doSomething(a)
			}
		}()
	}

	for i := 1; i < 50; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
