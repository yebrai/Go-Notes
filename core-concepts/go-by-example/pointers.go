package main

import "fmt"

// Function that takes an integer value and sets it to 0
func zeroval(ival int) {
	ival = 0
}

// Function that takes a pointer to an integer and sets the value at that address to 0
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// Initialize a variable i with the value 1
	i := 1
	fmt.Println("initial:", i) // Print the initial value of i

	// Call zeroval passing i by value
	zeroval(i)
	fmt.Println("zeroval:", i) // Print the value of i after calling zeroval

	// Call zeroptr passing the address of i (pointer to i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // Print the value of i after calling zeroptr

	// Print the memory address of i
	fmt.Println("pointer:", &i)
}
