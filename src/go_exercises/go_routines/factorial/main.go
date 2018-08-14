package main

import (
	"fmt"
)

func main() {
	c := factorial(4)

	for n := range c { // c is available as factorial is running, this also blocks main from exiting
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
