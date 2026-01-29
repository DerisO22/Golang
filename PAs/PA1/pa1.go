package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"
)

const FILE_TO_READ = "messages.txt";
const DATE_LAYOUT = "01/02/2006";

// reads and add to list
func readFromFile(messageArray *[]string, messageArrayLength *int) {
	// From Lab 3 Example of Reading File
	readFile, err := os.Open(FILE_TO_READ);
	
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}
	
	defer readFile.Close();

	scanner := bufio.NewScanner(readFile);
	fmt.Println("Reading from file:");

	for scanner.Scan() {
		*messageArray = append(*messageArray, scanner.Text());
		*messageArrayLength++;
	}
}

// generate rand num
func generateRandNumber(max_num int) int {
	// from lab 3 specifications
	max := big.NewInt(int64(max_num));

	tmp, err := rand.Int(rand.Reader, max);

	if err != nil {
		panic(err);
	}

	return int(tmp.Int64());
}

func calculateAge(inputBirthDate time.Time) int {
	currentTime := time.Now();

	// Age in years
	age := currentTime.Year() - inputBirthDate.Year();

	// "For example, if today is 1/20/2026 and the user is born on 1/21/2006, 
	// then the user is still 19 until the next day. On the other hand, if the 
	// user is born on 1/19/2006, the user is 20"
	if (currentTime.YearDay() < inputBirthDate.YearDay()){
		age--;
	}

	return age;
}

func main() {
	// using slice like the generateArray func from lab3 example
	var messages []string;
	var name string;
	var birthDate string;
	messageArrayLength := 0;

	// I could just return the new populated message slice
	// in readFromFile, but wanted to try out pointers/pass-by-ref
	readFromFile(&messages, &messageArrayLength);

	fmt.Println("\nI'm am the Almighty Psychic Fortune Teller!");
	fmt.Println("I know what's in your future!!!\n");

	// get info
	fmt.Print("What is your name?: ");
	fmt.Scan(&name);

	fmt.Print("When were you born (mm/dd/yyyy)?: ");
	fmt.Scan(&birthDate);

	formattedBirthDate, err := time.Parse(DATE_LAYOUT, birthDate);
	
	if err != nil {
		fmt.Println("Error formatting birth date: ", err);
		return;
	}

	userAge := calculateAge(formattedBirthDate);

	fmt.Printf("\n%s, you are %d years old\n", name, userAge);

}