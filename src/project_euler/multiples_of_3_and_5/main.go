package main

import "fmt"

func main() {
	var result = 0

	for num := 1; num < 1000; num++ {
		if num%3 == 0 || num%5 == 0 {
			result += num
		}
	}

	fmt.Println(result)
}
