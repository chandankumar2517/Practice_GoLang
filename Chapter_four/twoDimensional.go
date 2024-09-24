package main

import (
	"fmt"
)

func main() {

	var rows, cols int

	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)

	fmt.Print("Enter the number of cols: ")
	fmt.Scan(&cols)

	matrix := make([][]int, rows)

	if rows == cols {

		for i := 0; i < rows; i++ {

			matrix[i] = make([]int, cols)
			for j := 0; j < cols; j++ {
				fmt.Printf("Enter element [%d][%d]: ", i, j)
				fmt.Scan(&matrix[i][j])

			}
		}

		fmt.Println("You entered the following matrix:")
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				fmt.Print(matrix[i][j], " ")
			}
			fmt.Println()
		}

		// retrieve diagnal value from 2-D matrix

		mainDiag, offDiag := getDiagnoalValue(matrix)

		fmt.Println("Main diagonal:", mainDiag)
		fmt.Println("Off-diagonal:", offDiag)

		mainDiagResult := calculateAdditionMainDiag(mainDiag)
		offDiagResult := calculateAdditionoffDiag(offDiag)

		fmt.Println("Main diagonal Sum:", mainDiagResult)
		fmt.Println("Off-diagonal Sum :", offDiagResult)

		finalResult := getDifferenceBetweenDiagnol(mainDiagResult, offDiagResult)

		fmt.Println("Difference of sum :", finalResult)

	} else {
		fmt.Println("Please enter both rows and column equally")
	}
}

func getDifferenceBetweenDiagnol(mainDiagResult, offDiagResult int) int {
	return (mainDiagResult - offDiagResult)
}

func calculateAdditionMainDiag(mainDiag []int) int {

	sum := 0
	for _, value := range mainDiag {
		fmt.Println("calculateAdditionMainDiag i = ", value)
		sum += value
		fmt.Println("calculateAdditionMainDiag = ", sum)
	}

	return sum
}

func calculateAdditionoffDiag(offDiag []int) int {

	sum := 0
	for _, value := range offDiag {
		sum += value
		fmt.Println("calculateAdditionoffDiag = ", sum)
	}

	return sum
}

func getDiagnoalValue(matrix [][]int) ([]int, []int) {

	n := len(matrix)
	mainDiagnol := make([]int, n)
	offDiagonal := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				mainDiagnol[i] = matrix[i][j]
			}

			if i+j == n-1 {
				offDiagonal[i] = matrix[i][j]
			}
		}
	}

	return mainDiagnol, offDiagonal
}
