package gol_exercises

import "fmt"

func RunSlicesExcercise() {
	// 1)
	hobbies := [3]string{"Gaming", "Working Out", "Coding"}
	fmt.Printf("Hobbies Array: %v\n", hobbies)

	// 2)
	fmt.Printf("First Hobby: %s\n", hobbies[0])
	fmt.Printf("Everything Else: %v\n", hobbies[1:])

	// 3)
	firstAndSecond1 := hobbies[:2]
	fmt.Printf("First and Second elements slice: %v\n", firstAndSecond1)
	firstAndSecond2 := []string{hobbies[0], hobbies[1]}
	fmt.Printf("First and Second elements sliced another way: %v\n", firstAndSecond2)

	// 4)
	firstAndSecond1 = firstAndSecond1[1:3]
	fmt.Printf("First and Second slice re-sliced to have Second and Thid Elements: %v\n", firstAndSecond1)

	// 5)
	courseGoals := []string{"Learn Go", "Build a Go project"}
	fmt.Printf("Course Goals %v\n", courseGoals)

	// 6)
	courseGoals[1] = "Build an Awesome Go project"
	fmt.Printf("Course Goals %v\n", courseGoals)
	courseGoals = append(courseGoals, "Be proficient in Go programming")
	fmt.Printf("Course Goals %v\n", courseGoals)

	// 7)
	type Product struct {
		title string
		id    string
		price float64
	}

	products := []Product{
		{
			"Computer",
			"comp",
			1000.00,
		},
		{
			"Mouse",
			"mouse",
			20.00,
		},
		{
			"Keyboard",
			"kb",
			100.00,
		},
	}

	fmt.Printf("Available Products %v\n", products)

	products = append(products, Product{
		"Graphics Card",
		"gc",
		800.00,
	})

	fmt.Printf("Available Products %v\n", products)
}
