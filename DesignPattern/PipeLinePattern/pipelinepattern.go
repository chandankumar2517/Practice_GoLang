package pipelinepattern

import (
	"fmt"
)

/**
*  	The pipeline pattern is a design pattern that is particularly well-suited for concurrent and parallel programming in Golang.
*	It involves creating a series of stages, or pipelines, that process data sequentially or concurrently.
*	Each stage takes input from the previous stage, processes it, and passes the result to the next stage.
*  	This pattern is often used to implement data processing workflows, such as reading data from a source, transforming it,
*	and writing it to a destination
**/

func GenerateNumber(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func SquareNumbers(input <-chan int, out chan<- int) {
	for number := range input {
		out <- number * number
	}
	close(out)
}

func PrintNumbers(result <-chan int) {
	for numbers := range result {
		fmt.Println(numbers)
	}

}
