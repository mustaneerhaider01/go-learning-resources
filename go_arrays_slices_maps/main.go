package main

import "fmt"

func main() {
	// => make() simple makes memory management efficient by letting Go ahead of the fixed items that either
	// a slice or map is going to hold.
	// 1). if you know in advance how many items you want to add to slice, we can use make() to let Go know
	// that it should pre-allocate enough memory space for these elements
	// 2). Advantage is that Go will in-advance create behind the scene array with that much capcity so later
	// on when we add new values it doesn't have to allocate new memory space
	// 3). If we go beyond capacity value, only then Go will allocate new space for elements and essentially
	// creates new arrays
	// 4). This make() call tells Go to allocate enoug memory space for strigs array that stores five values
	// and then create that array with two empty slots so basically two empty strings
	userNames := make([]string, 2, 5)

	userNames[0] = "Joe"
	userNames[1] = "Ben"

	// Theses will be added to the end and will not fill that reserved slots
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	// 1). make() can also be used with maps to pre-allocate memory space for specified items during creation
	// 2). In case of map we can also pass one additional argument and that is the intended length for map
	// 3). Once we cross that capacity, only then Go will allocate new space for the items.
	courseRatings := make(map[string]float64)

	courseRatings["react"] = 4.9
	courseRatings["vue"] = 4.5
	courseRatings["nodejs"] = 4.7

	fmt.Println(courseRatings)
}
