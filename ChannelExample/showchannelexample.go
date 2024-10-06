package channelexample

import (
	"fmt"
	"sync"
)

func ShowChannel() {

	var wg sync.WaitGroup
	ch := make(chan int)

	// Start the workers
	for i := 0; i < 5; i++ {
		wg.Add(1)          //once per worker
		go worker(ch, &wg) // Each iteration starts a new worker goroutine
	}

	// Send values to the channel
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// Close the channel when all values have been sent
	close(ch)
	wg.Wait()

}

// multiple worker goroutines are created by the loop
func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Each worker listens for values from the channel until it's closed

	// The workers read from the channel (ch) until it is closed.
	// This pattern ensures safe communication without race conditions or the need for explicit locking
	for result := range ch {
		fmt.Println(result)
	}
}
