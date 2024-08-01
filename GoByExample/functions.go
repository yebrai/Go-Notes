package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func multipleReturnValues(a int, b bool) (int, bool) {
	return a, b
}

func multipleReturnValuesNoParam() (int, bool) {
	// Retornamos un valor entero y un valor booleano
	return 42, true
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := multipleReturnValuesNoParam()
	fmt.Println("multipleReturnValuesNoParam", a, b)

	resultInt, resultBool := multipleReturnValues(5, true)
	fmt.Printf("ResultInt: %d, ResultBool: %t\n", resultInt, resultBool)

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
