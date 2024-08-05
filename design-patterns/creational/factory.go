package creational

import "fmt"

// Animal interface defines a common behavior (Speak) for all animals.
type Animal interface {
	Speak() string
}

// Dog struct implements the Animal interface.
type Dog struct{}

// Speak method for Dog returns a sound specific to dogs.
func (d Dog) Speak() string { return "Woof!" }

// Cat struct implements the Animal interface.
type Cat struct{}

// Speak method for Cat returns a sound specific to cats.
func (c Cat) Speak() string { return "Meow!" }

// AnimalFactory is a factory function that creates and returns an Animal object
// based on the provided animalType string.
func AnimalFactory(animalType string) Animal {
	if animalType == "dog" {
		return Dog{}
	} else if animalType == "cat" {
		return Cat{}
	}
	return nil // Return nil if the animal type is unknown.
}

func main() {
	// Use the AnimalFactory to create an Animal object.
	animal := AnimalFactory("dog")
	if animal != nil {
		// Call the Speak method of the created animal.
		fmt.Println(animal.Speak())
	} else {
		fmt.Println("Unknown animal type")
	}
}
