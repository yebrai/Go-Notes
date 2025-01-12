package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("halo, Go!")

	var myString string = "Hello, Go!"
	fmt.Println(myString)

	var myInt int = 7
	fmt.Println(myInt)

	fmt.Printf("%s %d", myString, myInt)
	fmt.Println(myString, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	//Array

	var myArray [3]int = [3]int{1, 2, 3}
	fmt.Println(myArray)

	var myArray2 [3]int
	myArray[0] = 1
	myArray2[0] = 2

	fmt.Println(myArray2)

	// Map

	myMap := make(map[string]int)
	myMap["Ivan"] = 33
	myMap["Ivan2"] = 33

	fmt.Println(myMap)
	fmt.Println(myMap["Ivan"])

	// List

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println("List", myList.Back().Value.(int))

	// Bucles

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray {
		fmt.Println(index, value)
	}

	//Funcion

	fmt.Println(myFunction())

	//Structure
	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Ivan", 33}
	fmt.Println(myStruct)

}

func myFunction() string {
	return "My function"
}
