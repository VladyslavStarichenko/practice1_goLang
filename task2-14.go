package main

import "fmt"

func isScaleneTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return a+b > c && a+c > b && b+c > a
}

func printResultScaleneTriangle(a, b, c float64) {
	fmt.Printf("Task 2 - variant 14\nTest sides: %.2f, %.2f, %.2f\nResult: %v\n\n", a, b, c, isScaleneTriangle(a, b, c))
}

func testScaleneTriangle() {
	// Test case with a scalene triangle
	a, b, c := 3.5, 4.2, 5.6
	printResultScaleneTriangle(a, b, c)

	// Test case with a degenerate triangle
	a, b, c = 3, 4, 7
	printResultScaleneTriangle(a, b, c)

	// Test case with zero or negative sides
	a, b, c = 2, 0, 5
	printResultScaleneTriangle(a, b, c)
}
