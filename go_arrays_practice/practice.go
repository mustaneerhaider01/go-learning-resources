package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"Coding", "Sports", "Cooking"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3
	mainHobbies := hobbies[:2]
	// mainHobbies := hobbies[0:2] -> other way
	fmt.Println(mainHobbies)

	// 4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5
	courseGoals := []string{"Learn Go", "Build REST APIs"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "Basics of Go"
	courseGoals = append(courseGoals, "Make RESTful Web services")
	fmt.Println(courseGoals)

	// 7
	products := []Product{
		{
			id:    1,
			title: "A Book",
			price: 59.99,
		},
		{
			id:    2,
			title: "A Mobile",
			price: 89.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		3,
		"A Carpet",
		49.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
