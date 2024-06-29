package main

import (
	"errors"
	"fmt"
	"time"
)

// Main function - entry point of the program
func main() {
	// String manipulation
	str := createString("Hello", "Go")
	fmt.Println("Created string:", str)

	// Sum of two numbers
	sumResult := sum(10, 20)
	fmt.Println("Sum of 10 and 20:", sumResult)

	// Array manipulation
	array := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", array)
	modifiedArray := modifyArray(array)
	fmt.Println("Modified array:", modifiedArray)

	// Error handling
	err := checkEvenNumber(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Number is even")
	}

	// Concurrent function execution
	fmt.Println("Starting concurrent execution")
	go concurrentFunction(1)
	go concurrentFunction(2)

	// Adding a sleep to allow goroutines to complete
	time.Sleep(1 * time.Second)
}

// Function to create a string by concatenation
func createString(part1 string, part2 string) string {
	return part1 + " " + part2
}

// Function to sum two numbers
func sum(a int, b int) int {
	return a + b
}

// Function to modify an array by adding 10 to each element
func modifyArray(arr []int) []int {
	for i := range arr {
		arr[i] += 10
	}
	return arr
}

// Function to check if a number is even and return an error if it's not
func checkEvenNumber(num int) error {
	if num%2 != 0 {
		return errors.New("the number is not even")
	}
	return nil
}

// Concurrent function that prints a message with a delay
func concurrentFunction(id int) {
	fmt.Printf("Concurrent function %d started\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Concurrent function %d finished\n", id)
}
