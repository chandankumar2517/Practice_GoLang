package main

import (
	"crypto/sha256"
	"fmt"
	"time"

	binarysearch "github.com/chandankumar2517/BinarySearch"
	channelexample "github.com/chandankumar2517/ChannelExample"
	checkFactoryDesignExample "github.com/chandankumar2517/DesignPattern/FactoryPattern"
	pipelinedesignpattern "github.com/chandankumar2517/DesignPattern/PipeLinePattern"
	checkSingletonExample "github.com/chandankumar2517/DesignPattern/SingletonPattern"
	goroutinecommunication "github.com/chandankumar2517/GoRoutineCommunicationExample"
	maxSubArray "github.com/chandankumar2517/MaximumSubArray"
	reversedString "github.com/chandankumar2517/ReverseGivenString"
	sha "github.com/chandankumar2517/ShaAlgorithm"
	twosum "github.com/chandankumar2517/TwoSum"
	compositionexample "github.com/chandankumar2517/compositionExample"
)

func main() {

	learnReverseGivenString()

	learGoRoutineCommunication()

	learWorkerPoolCommunication()

	//learnSHAProblem()

	//learnTwoSumProblem()

	//learnBinarySearch()

	//learnMaximumSubArraySum()

	//learnComposition()

	//learnChannelCommunication()

	//learnSingletonDesingPattern()

	//learnFactoryDesingPattern()

	//learnPipeLineDesignPattern()

}

func learnReverseGivenString() {
	s1 := "Chandan"
	fmt.Println("Provided String to be reversed ", s1)
	reversedString := reversedString.ReverseGivenString(s1)

	fmt.Println("Reversed String is ", reversedString)
}

func learGoRoutineCommunication() {
	goroutinecommunication.LearnCommunication()
	print("\n")
}

func learWorkerPoolCommunication() {
	goroutinecommunication.WorkerPool()
}

/**This function helps you to
* convert two given string and find out
* total count of different element
**/
func learnSHAProblem() {
	fmt.Println("******* SHA-256 Problem begain from here ******* ")

	// SHA 256 Algorithm
	hash1 := sha256.Sum256([]byte("a"))
	hash2 := sha256.Sum256([]byte("A"))

	fmt.Printf("x = %x \n", hash1)
	fmt.Printf("y = %x \n", hash2)

	// now compare and count number of bits in two SHA

	diffbits := sha.CountDifferentBits(hash1, hash2)

	fmt.Println("Number of different item count is ", diffbits)
}

// TWO Sum Problem
func learnTwoSumProblem() {
	fmt.Println("\n******* Two Sum Problem begain from here ******* ")
	arry := []int{1, 2, 3, 4, 5, 6}
	target := 8

	indices := twosum.TwoSum(arry, target)
	fmt.Printf("For a given array %d\ntarget %d \nIndices are %v \n", arry, target, indices)
}

// understand and learn how binary search tree works
// Binary search Problem
func learnBinarySearch() {

	bst := binarysearch.BST{}

	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(12)
	bst.Insert(9)
	bst.Insert(11)

	fmt.Println("\n\n******* Binary Serch begain from here *******")

	fmt.Println("In-order traversal of BST:")
	bst.InOrder() // This will print all values in sorted orde

	fmt.Println("Item found or not = ", bst.Search(5))

	bst.Delete(5)

	fmt.Println("Item found (5) after deletion:", bst.Search(5)) // Output: false

	fmt.Println("In-order traversal of After delete BST:")

	bst.InOrder() // This will print all values in sorted order
}

// this function helps to find maximum sub array when added together it should give
// maximum value from a given array
func learnMaximumSubArraySum() {
	fmt.Println("\n\n******* Maximum SubArray Problem begain from here *******")

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	maxSum, subarray := maxSubArray.GetMaximumSubArray(nums)
	fmt.Println("Maximum Subarray Sum:", maxSum)
	fmt.Println("Subarray:", subarray)
}

// composition example
func learnComposition() {
	newGame := compositionexample.Game{
		Name:    "TikTok",
		Price:   "500",
		Details: compositionexample.Details{Genre: "Action", GeneralRating: "4.5", Reviews: "Good game"},
	}

	newGame.ShowGameDetails()
}

// simple example for goroutines communicate using channels
func learnChannelCommunication() {
	channelexample.ShowChannel()
}

func learnSingletonDesingPattern() {
	fmt.Println("*** Singleton Desing Pattern **** ")
	S1 := checkSingletonExample.GetInstance()
	S2 := checkSingletonExample.GetInstance()

	result := S1 == S2
	fmt.Println("Singelton instace are same or not = ", result)

}

func learnFactoryDesingPattern() {
	fmt.Println("*** Factory Desing Pattern **** ")

	facotry := checkFactoryDesignExample.ShapeFactory{}

	rectangle := facotry.CreateShape("Reactangle")
	square := facotry.CreateShape("Square")

	rectangle.Draw()
	square.Draw()
}

func learnPipeLineDesignPattern() {
	fmt.Println(" **** Pipe Line Design Pattern **** ")

	number := make(chan int)
	square := make(chan int)

	go pipelinedesignpattern.GenerateNumber(number)
	go pipelinedesignpattern.SquareNumbers(number, square)
	go pipelinedesignpattern.PrintNumbers(square)

	time.Sleep(time.Second)
}
