package main

import "fmt"

func main() {
	var a [3]int // array by default have 0
	var b []int  // by defualt it will not have anything if we don't mention
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(len(a))
	fmt.Println(len(b))

	// print both indices and element
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print only element
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

	// slice related program

	c := make([]int, 4)

	d := append(c, 4)

	c[3] = 7

	fmt.Println(c, d)

	/*r := [...]int{3: -1}

	s := []int{3: -1}

	fmt.Println(r)
	fmt.Println(len(r))
	fmt.Println(s)
	fmt.Println(len(s))*/

	r1 := []int{3: -1}
	fmt.Println(r1, len(r1), cap(r1)) // Output: [-1] 1 4

	r2 := [...]int{3: -1}
	fmt.Println(r2, len(r2), cap(r2)) // Output: [0 0 0 -1] 4 4
}
