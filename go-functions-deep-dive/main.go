package main

import "fmt"

// We have simply assigned a type alias to a function type.
// TIP: Use type aliases for more longer function types
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumers := []int{5, 1, 2}

	doubled := transformNumbers(&numbers, double) // We pass the name of the function as value
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	// Both transformFn1, transformFn2 stores functions as values and therefore can also be passed as values around.
	transformerFn1 := getTransformerFn(&numbers)
	transformerFn2 := getTransformerFn(&moreNumers)

	tranformedNums := transformNumbers(&numbers, transformerFn1)
	moreTranformedNums := transformNumbers(&moreNumers, transformerFn2)

	fmt.Println(tranformedNums)
	fmt.Println(moreTranformedNums)
}

// Functions in GO are first class values, so we can pass them as normal values and therefore also accept them
// as paramters inside other functions.
func transformNumbers(nums *[]int, transform transformFn) []int {
	dNums := []int{}

	// With slices, we have to manually dereference the pointer unlike structs where GO does it for us.
	for _, val := range *nums {
		dNums = append(dNums, transform(val))
	}

	return dNums
}

// Functions can also return other functions as values.
func getTransformerFn(nums *[]int) transformFn {
	if (*nums)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
