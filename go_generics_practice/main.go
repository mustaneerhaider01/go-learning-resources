package main

import "fmt"

// We have assigned type-aliased to built-in 'string' type that allows us to define methods
// for the custom type
type str string

func (text str) log() {
	fmt.Println("Result:", text)
}

func main() {
	result := add(5, 10)

	resultLog := str(fmt.Sprint(result))
	resultLog.log()
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
