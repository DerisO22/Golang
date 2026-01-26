package main

import (
	"fmt"
	"log"
)

func main() {
	// Lab 2 Leap Year Check - Deris O
	var inputYear int;
	totalAttempts := 0;

	// user input
	for inputYear <= 0 {
		if totalAttempts > 0 {
			fmt.Print(("Invalid Input! Try Again.\n"))
		}

		fmt.Print("\nEnter a year: ");
		// From code 2
		_, err := fmt.Scanln(&inputYear);

		if err != nil {
			log.Fatal(err);
		}

		totalAttempts++;
	}

	//check leap
	if inputYear % 4 == 0 && (inputYear % 100 != 0 || inputYear % 400 == 0) {
		fmt.Printf("\n%d is a valid leap year", inputYear);
	} else {
		fmt.Printf("\n%d is not a valid leap year", inputYear);
	}
}