package main

import "fmt"

func PrintType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("INT value: %v \n", v)
	case string:
		fmt.Printf("String value: %v \n", v)
	case bool:
		fmt.Printf("Bool value: %v \n", v)
	default:
		fmt.Printf("The value was not recognized %v \n", v)
	}
}

func main() {
	PrintType("Hello")
	PrintType(1)
	PrintType(true)
	PrintType(10.20)
}
