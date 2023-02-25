package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(n int) bool {
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

func printResultPrefectSquare(num1 int, num2 int) {
	fmt.Println("Task 1 - variant 18")
	fmt.Printf("%d is a perfect square: %v\n", num1, isPerfectSquare(num1))
	fmt.Printf("%d is a perfect square: %v\n", num2, isPerfectSquare(num2))
}
