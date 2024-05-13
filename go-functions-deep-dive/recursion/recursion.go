package recursion

import "fmt"

func main() {
	fact := factorial(7)
	fmt.Println(fact)
}

// Recursion is when function call itself.
// The function can't finish it's execution unless the inner function(s) it calls return their values
// and finish their execution.
// When a function recursively call itself, there should be a terminating condition to end the function
// execution else this will create infinite loop.
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
