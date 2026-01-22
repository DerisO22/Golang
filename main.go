package main

import (
	"fmt"
	"math"
)

type Solution struct {}

func (s *Solution) checkPalindrome(number int) bool {
	// abs value of number - base case
	math.Abs(float64(number));

	originalNum := number

	reverseNum := 0

	for number != 0 {
		// Get the last digit and then add to a temporary
		lastDigit := number % 10
		reverseNum = reverseNum * 10 + lastDigit
		number = number / 10
	}

	return reverseNum == originalNum
}	

func main() {
	// Leet-Code Problem - isPalindrome
	solution_obj := &Solution{}

	isPalindrome := solution_obj.checkPalindrome(202)
	
	fmt.Println("Is a palindrome?:", isPalindrome)
}