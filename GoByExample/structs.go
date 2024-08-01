package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

var arrayExample = [3]int{1, 2, 3}
var mapExample = map[string]int{"one": 1, "two": 2}
var structExample = person{"Bob", 20}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(reflect.TypeOf(person{"Bob", 20}))

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	fmt.Println("Type of arrayExample:", reflect.TypeOf(arrayExample))
	fmt.Println("Type of mapExample:", reflect.TypeOf(mapExample))
	fmt.Println("Type of structExample:", reflect.TypeOf(structExample))
}
