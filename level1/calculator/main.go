package main

import (
	"errors"
	"fmt"
	"log"
)

func Add(num1, num2 int) int {
	return num1 + num2
}

func Subtract(num1, num2 int) int {
	return num1 - num2
}

func Multiply(num1, num2 int) int {
	return num1 * num2
}

func Divide(num1, num2 int) (float32, error) {
	if num2 <= 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return float32(num1 / num2), nil
}

func main() {
	fmt.Println(Add(2, 3))
	fmt.Println(Subtract(2, 3))
	fmt.Println(Multiply(2, 3))
	result, err := Divide(6, 0)

	if err != nil {
		defer log.Panic(err.Error())
	} else {
		fmt.Println(result)
	}

	fmt.Println(Divide(6, 2))
}
