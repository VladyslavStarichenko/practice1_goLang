package main

import "fmt"

func printPythagoreanNumbersResult() {
	fmt.Println("Task 4 - variant 14")
	fmt.Println("Test case - 20")
	printPythagoreanNumbers(20)
}

func printPythagoreanNumbers(n int) {
	for a := 1; a < n; a++ {
		for b := a + 1; b < n; b++ {
			for c := b + 1; c < n; c++ {
				if a*a+b*b == c*c {
					fmt.Printf("%d, %d, %d\n", a, b, c)
				}
			}
		}
	}
}
