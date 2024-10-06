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
