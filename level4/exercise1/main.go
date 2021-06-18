package main

import "fmt"

var (
	msg = func() {
		fmt.Println("Hello world")
	}
)

func main() {
	go msg()
	fmt.Println("main function")
}
