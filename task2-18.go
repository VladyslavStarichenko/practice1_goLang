package main

import (
	"fmt"
	"strconv"
)

func countPositiveNumbers(a, b, c int) int {
	counter := 0
	increaseCounter(checkPositive(a), counter)
	increaseCounter(checkPositive(b), counter)
	increaseCounter(checkPositive(c), counter)
	return counter
}

func checkPositive(number int) bool {
	return number > 0
}

func increaseCounter(isPositive bool, count int) {
	if isPositive {
		count++
	}
}

func printCountPositiveNumbers() {
	a := 5
	b := -5
	c := -25
	fmt.Println("Task 2 - variant 18")
	fmt.Println("Test digit: " + strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c))
	fmt.Println("Result: " + strconv.Itoa(countPositiveNumbers(a, b, c)))

}
