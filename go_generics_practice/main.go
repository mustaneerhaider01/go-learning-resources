package main

import "fmt"

// We have assigned type-aliased to built-in 'string' type that allows us to define methods
// for the custom type
type str string

func (text str) log() {
	fmt.Println("Result:", text)
}

func main() {
	// Since both param types is int so Go correclty infers the return type here as 'int'
	result := add(5, 10)

	// first convert result to string as direct coversion to str() wil yield a 'rune'
	// Then covert string to str()
	resultLog := str(fmt.Sprint(result))
	resultLog.log()
}

// T refers to type placeholder and we have list out specific type that can be filled in for T
// The type for T will be set when we call this add()
// T basically is a union of these three types and will not accept value of any other type, so both
// paramters to this add() should of these listed types.
func add[T int | float64 | string](a, b T) T {
	return a + b
}
