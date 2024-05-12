package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// With closures, values access from outer scope are locked in when the anonynmous function is created.
	// When we call createTransform with 2 as an argument then 2 will be locked in and then later on when this double function is executed
	// the number will always be multiplied with 2.
	double := createTransformer(2)
	triple := createTransformer(3)

	// If you have a function that is only used in one palce in the app, you can use anonymous functions.
	// Anonymous functions are created JIT.
	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Factory function...
// We can return anonymous functions as well since functiona themselves are returnable form other functions.
// We are locking in the value of factor parameter since anonymous function create closures in GO.
func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}
