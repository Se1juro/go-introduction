package main

import "fmt"

func main() {
	var greeting string = "Hello world"
	greet := "Hello world without string"

	fmt.Println(greeting)
	fmt.Println(greet)

	// Arrays
	var ages [2]int

	// ages[0] = 6
	ages[1] = 4

	// Print [0, 4]
	fmt.Println(ages)

	// Slices
	numbers := []int{5, 9, 10}
	numbers = append(numbers, 8, 11, 3)
	fmt.Println(numbers)

	var numberSlices = []int{}
	fmt.Println(numberSlices)
	numberSlices = append(numberSlices, 10, 30, 50)
	fmt.Println(numberSlices)

	numbers = numbers[0:5]
	fmt.Println(numbers)

	// Loops

	// Infinite loop
	/* for {
		fmt.Println("Not finish")
	} */

	//len(numbers) Length of slices
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("===Simple for loop slices===")
	// Simple for slices
	for i, number := range numbers {
		fmt.Println("Index:", i, "Value:", number)
	}
	fmt.Println("===Without indexes===")
	// Without index
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Value:", number)
		}
	}

	fmt.Println("===FRUITS===")
	fruits := []string{"Apple", "Banana", "Watermelon", "Melon"}

	for _, fruit := range fruits {
		if fruit != "Banana" {
			continue
		}
		fmt.Println("Fruit: ", fruit)
	}

	fmt.Println("===LAST LETTER===")

	for _, fruit := range fruits {
		lastLetter := fruit[len(fruit)-1:]

		if lastLetter == "a" {
			continue
		}
		fmt.Println("Fruit: ", fruit)
	}

}
