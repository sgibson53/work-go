package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, say whatever you want:")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello, there.")
	fmt.Print("You said: ")
	fmt.Println(text)
}
