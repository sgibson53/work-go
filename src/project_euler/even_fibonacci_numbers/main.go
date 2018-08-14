package main

import "fmt"

func main() {
	var fibs = generateFibonacciNumbers()
	fmt.Println(fibs)

	var answer int

	for _, num := range fibs {
		if num%2 == 0 {
			answer += num
		}
	}

	fmt.Println(answer)
}

func generateFibonacciNumbers() []int {
	result := []int{0, 1}
	var a, b = 0, 1
	var newNum int
	for {
		newNum = a + b
		a = b
		b = newNum

		if newNum < 4000000 {
			result = append(result, newNum)
		} else {
			return result
		}
	}
}
