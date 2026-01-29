package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const FILE_TO_READ = "messages.txt"

// reads and add to list
func readFromFile(messageArray *[]string) {
	// From Lab 3 Example of Reading File
	readFile, err := os.Open(FILE_TO_READ);
	
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}
	
	defer readFile.Close();

	scanner := bufio.NewScanner(readFile);
	fmt.Println("Reading from file:");

	i := 0;

	for scanner.Scan() {
		*messageArray = append(*messageArray, scanner.Text());
		fmt.Println((*messageArray)[i]);
		i++;
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

func main() {
	// using slice like the generateArray func from lab3 example
	var messages []string;

	// I could just return the new populated message slice
	// in readFromFile, but wanted to try out pointers/pass-by-ref
	readFromFile(&messages);


}