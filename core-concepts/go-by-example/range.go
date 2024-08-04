package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 || num == 4 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for _, v := range kvs {
		fmt.Println("value:", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	str := "hello"
	for i, c := range str {
		fmt.Printf("Index: %d, Character: %c\n", i, c)
	}
}
