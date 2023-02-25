package main

import "fmt"

func printFindNumbers() {
	fmt.Println("Task 4 - variant 18")
	fmt.Println("Test case - 100, 25")
	fmt.Println(findNumbers(100, 25))
}

func findNumbers(n, m int) []int {
	var result []int
	for i := 1; i < n; i++ {
		sum := 0
		num := i
		for num > 0 {
			digit := num % 10
			sum += digit
			num /= 10
		}
		if sum*sum == m {
			result = append(result, i)
		}
	}
	return result
}
