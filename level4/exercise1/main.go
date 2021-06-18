package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello world")
	}()

	fmt.Println("main function")
}
