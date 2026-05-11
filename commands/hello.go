package commands

import "fmt"

func Hello(args []string) {
	fmt.Println("Hello, World!")
}

func init() {
	Register("hello", Hello)
}
