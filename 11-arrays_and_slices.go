package main

import "fmt"

func main() {
	// Golang uses zero-based indexes for array and slices

	// Arrays have a defined length and are immutable
	var defaults [4]int
	fmt.Println("Array: ", defaults)
	fmt.Println("Array length: ", len(defaults))
	fmt.Println("Array capacity: ", cap(defaults))

	defaults[0] = 1
	defaults[1] = 2
	defaults[2] = 3
	defaults[3] = 4
	fmt.Println("Array: ", defaults)
	fmt.Println("Array length: ", len(defaults))
	fmt.Println("Array capacity: ", cap(defaults))

	// Slices do not have a defined length
	evenNumbers := []int{2, 4, 6, 8, 10, 12, 14, 16}
	fmt.Println("Slice: ", evenNumbers)
	fmt.Println("Slice length: ", len(evenNumbers))
	fmt.Println("Slice capacity: ", cap(evenNumbers))

	// Get value at index
	fmt.Println(evenNumbers[0])

	// Get values from index 1 to before index 3
	fmt.Println(evenNumbers[1:3]) // Value at index 3 is excluded

	// Get values from first index before index 5
	fmt.Println(evenNumbers[:5]) // Value at index 5 is excluded

	// Get values from index 2 to last index
	fmt.Println(evenNumbers[2:])

	// Adding one new element
	evenNumbers = append(evenNumbers, 18)
	fmt.Println("Slice: ", evenNumbers)

	// Adding new elements using the spread operator
	moreEvenNumbers := []int{20, 22, 24}
	evenNumbers = append(evenNumbers, moreEvenNumbers...)
	fmt.Println("Slice: ", evenNumbers)
}
