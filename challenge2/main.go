package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidCreditCardNumber(cardNumber string) bool {
	// Remove any hyphens
	cardNumberWithoutHyphen := strings.ReplaceAll(cardNumber, "-", "")

	// Check if the first digit is 4, 5, or 6
	firstDigit, _ := strconv.Atoi(string(cardNumberWithoutHyphen[0]))
	if firstDigit != 4 && firstDigit != 5 && firstDigit != 6 {
		return false
	}

	// Check if the length is exactly 16 digits
	if len(cardNumberWithoutHyphen) != 16 {
		return false
	}

	// Check if all characters are digits
	for _, digit := range cardNumberWithoutHyphen {
		if digit < '0' || digit > '9' {
			return false
		}
	}

	// If the card number contains hyphens, check if digits are grouped in sets of 4
	if strings.Contains(cardNumber, "-") && validGrouping(cardNumber) {
		return true
	}

	return false
}

func validGrouping(cardNumber string) bool {
	digits := strings.Split(cardNumber, "-")
	if len(digits) != 4 {
		return false
	}
	for _, group := range digits {
		if len(group) != 4 {
			return false
		}
	}
	return true
}

func main() {
	// Read the number of test cases
	var n int
	fmt.Println("Enter the number of test cases:")
	fmt.Scanln(&n)

	// Read all the test cases
	fmt.Println("Enter the credit card numbers:")
	var testCases []string
	for i := 0; i < n; i++ {
		var cardNumber string
		fmt.Scanln(&cardNumber)
		testCases = append(testCases, cardNumber)
	}

	// Validate each test case
	for _, cardNumber := range testCases {
		if isValidCreditCardNumber(cardNumber) {
			fmt.Println(cardNumber, "is a valid credit card number.")
		} else {
			fmt.Println(cardNumber, "is an invalid credit card number.")
		}
	}
}
