package main

import (
	"fmt"
	"strconv"
)

func printQuantityOfDigitsResult() {
	fmt.Println("Task 3 - variant 18")
	fmt.Println("Test case - -100000000000012z")
	printQuantityOfDigits("-100000000000012z")
}

func printQuantityOfDigits(inputString string) {
	num, err := strconv.Atoi(inputString)

	if err != nil {
		fmt.Println("You entered not a number.")
		return
	}

	if num < 1 {
		fmt.Println("Natural number is a number, which is greater than zero.")
		return
	}

	formattedNum := strconv.Itoa(num)

	formattedNumLength := len(formattedNum)

	inputStringLength := len(inputString)

	if formattedNumLength != inputStringLength {
		fmt.Println("You've entered the number with leading zeroes.")
		fmt.Println("We need to format it. Please, wait a second...")
		fmt.Println("Here is the length with formatted number: " + strconv.Itoa(formattedNumLength))
		return
	}

	fmt.Println("Here is the length of the entered number: " + strconv.Itoa(inputStringLength))
}
