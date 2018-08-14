package main

import "fmt"

func main() {
	fmt.Println("Function #1, num divided by two and is it even:")
	var numDividedByTwo, isNumEven = first(7)
	fmt.Println(fmt.Sprintf("(%d, %t)", numDividedByTwo, isNumEven))

	fmt.Println("Function #2, highest num in list")
	fmt.Println(highestNumInVariadicParam(3, 7, 9, 4, 3, -9, 0))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(highestNumInVariadicParam(aSlice...))
	fmt.Println(highestNumInVariadicParam())
}

/* Write a funciton which takes an integer. The function will have two
returns. The first return should be the argument divided by 2.
The second return should be a bool that let's us know whether or not the argument was even. */
func first(a int) (int, bool) {
	var isEven = false

	if a%2 == 0 {
		isEven = true
	}

	var result = a / 2

	return result, isEven
}

/* Write a function with one variadic parameter taht finds the greatest number in a list of numbers */
func highestNumInVariadicParam(nums ...int) int {
	var highest int
	for _, num := range nums {
		if highest == 0 || num > highest {
			highest = num
		}
	}

	return highest
}
