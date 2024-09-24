package shaalgorithm

import (
	"fmt"
)

// Write a function that counts the number of bits that are different in two SHA256 hashes.

func CountDifferentBits(hash1, hash2 [32]byte) int {

	differentBits := 0

	for i, _ := range hash1 {
		xorResult := hash1[i] ^ hash2[i]

		fmt.Println("XOR Operation ", xorResult)

		differentBits += countBits(xorResult)
	}

	return differentBits

}

func countBits(b byte) int {
	count := 0

	for b > 0 {
		count += int(b & 1) // Check if the least significant bit is 1

		b >>= 1 // Right shift to check the next bit
	}

	return count
}
