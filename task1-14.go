package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func checkDigitsUnique(digit int) bool {
	stringValue := strconv.Itoa(digit)
	digitMap := make(map[string]int)
	for i := 0; i < len(stringValue); i++ {
		digitMap[string(stringValue[i])] = int(stringValue[i])
	}

	return len(stringValue) == len(digitMap)
}

func printResultUniqueDigits() {
	digit2Check := generateRandomNumber()
	fmt.Println("Task 1 - variant 14")
	fmt.Println("Test digit: " + strconv.Itoa(digit2Check))
	fmt.Println("Result: " + strconv.FormatBool(checkDigitsUnique(digit2Check)) + "\n")
}

func generateRandomNumber() int {
	return rand.Intn(9000)
}
