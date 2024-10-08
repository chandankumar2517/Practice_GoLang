package goroutinecommunicationexample

import (
	"fmt"
	"sync"
)

func LearnCommunication() {
	var wg sync.WaitGroup

	ch := make(chan int)

	// Create a goroutine to send a value
	go func() {
		defer wg.Done()
		fmt.Println("Sending value...")
		ch <- 42

	}()

	go func() {

		defer wg.Done()
		value := <-ch
		fmt.Println("Received value from channel ", value)
	}()

	wg.Add(2)
	wg.Wait()

}

func WorkerPool() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(ch, &wg)
		ch <- i
	}

	close(ch)
	wg.Wait()

}

func worker(value <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	result := <-value

	fmt.Println(result)

}
