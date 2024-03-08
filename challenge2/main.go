package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidCreditCardNumber(cardNumber string) bool {
	// Remove any hyphens
	cardNumberWithoutHyphen := strings.ReplaceAll(cardNumber, "-", "")

	// Check if the length is exactly 16 digits
	if len(cardNumberWithoutHyphen) != 16 {
		return false
	}

	// Check if the first digit is 4, 5, or 6
	firstDigit, _ := strconv.Atoi(string(cardNumberWithoutHyphen[0]))
	if firstDigit != 4 && firstDigit != 5 && firstDigit != 6 {
		return false
	}

	// Check if all characters are digits
	for _, digit := range cardNumberWithoutHyphen {
		if digit < '0' || digit > '9' {
			return false
		}
	}

	// If the card number contains hyphens, check if digits are grouped in sets of 4
	if strings.Contains(cardNumber, "-") && !validGrouping(cardNumber) {
		return false
	}

	// Check for consecutive repeated digits (not exceeding 3)
	repeated := 0
	for i := 0; i < len(cardNumberWithoutHyphen)-1; i++ {
		if cardNumberWithoutHyphen[i] == cardNumberWithoutHyphen[i+1] {
			repeated++
			if repeated >= 3 {
				return false
			}
		} else {
			repeated = 0
		}
	}

	return true
}

func validGrouping(cardNumber string) bool {
	digits := strings.Split(cardNumber, "-")
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

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func isValidCreditCardNumber(cardNumber string) bool {
// 	// Remove any hyphens
// 	cardNumberWithoutHyphen := strings.ReplaceAll(cardNumber, "-", "")

// 	// Check if the length is exactly 16 digits
// 	if len(cardNumberWithoutHyphen) != 16 {
// 		return false
// 	}

// 	// Check if the first digit is 4, 5, or 6
// 	firstDigit, _ := strconv.Atoi(string(cardNumberWithoutHyphen[0]))
// 	if firstDigit != 4 && firstDigit != 5 && firstDigit != 6 {
// 		return false
// 	}

// 	// Check if all characters are digits
// 	for _, digit := range cardNumberWithoutHyphen {
// 		if digit < '0' || digit > '9' {
// 			return false
// 		}
// 	}

// 	// If the card number contains hyphens, check if digits are grouped in sets of 4
// 	if strings.Contains(cardNumber, "-") && !validGrouping(cardNumber) {
// 		return false
// 	}

// 	// Check for consecutive repeated digits (not exceeding 3)
// 	repeated := 0
// 	for i := 0; i < len(cardNumberWithoutHyphen)-1; i++ {
// 		if cardNumberWithoutHyphen[i] == cardNumberWithoutHyphen[i+1] {
// 			repeated++
// 			if repeated >= 3 {
// 				return false
// 			}
// 		} else {
// 			repeated = 0
// 		}
// 	}

// 	return true
// }

// func validGrouping(cardNumber string) bool {
// 	digits := strings.Split(cardNumber, "-")
// 	for _, group := range digits {
// 		if len(group) != 4 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	// Example test cases
// 	// cardNumbers := []string{"4253-6258-7961-5786", "4424424424442444", "5122-2368-7954-3214", "4253-6258-7961-57867", "4424444424442444", "5122-2368-7954 - 3214", "6123-4567-8912-3456", "51-67-8912-3456", "4111111111111111"}
// 	cardNumbers := []string{"4123456789123456", "5123-4567-8912-3456", "61234-567-8912-3456", "4123356789123456", "5133-3367-8912-3456", "5123 - 3567 - 8912 - 3456	"}

// 	// Check each card number
// 	for _, cardNumber := range cardNumbers {
// 		if isValidCreditCardNumber(cardNumber) {
// 			fmt.Println(cardNumber, "is a valid credit card number.")
// 		} else {
// 			fmt.Println(cardNumber, "is an invalid credit card number.")
// 		}

// 	}
// }

// cardNumbers := []string{"4123456789123456", "5123-4567-8912-3456", "61234-567-8912-3456"}
