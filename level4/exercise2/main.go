package main

import (
	"fmt"
	"time"
)

func goroutine() {
	func() {
		fmt.Println("1")
	}()
}

func main() {
	go goroutine()
	fmt.Println("0")
	time.Sleep(2 * time.Second)
}
