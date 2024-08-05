package creational

import (
	"fmt"
	"sync"
)

// Singleton represents the single instance we want to create.
type Singleton struct {
	Message string
}

// instance is the only instance of Singleton.
var instance *Singleton
var once sync.Once

// GetInstance returns the single instance of Singleton, creating it if necessary.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Message: "Hello from the Singleton!"}
	})
	return instance
}

func main() {
	// Get the Singleton instance for the first time
	s1 := GetInstance()
	fmt.Println(s1.Message)

	// Get the Singleton instance again
	s2 := GetInstance()
	fmt.Println(s2.Message)

	// Check if both variables point to the same instance
	fmt.Printf("s1 and s2 are the same instance: %v\n", s1 == s2)
}
