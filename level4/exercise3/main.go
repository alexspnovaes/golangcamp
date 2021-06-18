package main

import (
	"fmt"
)

func addValue(c chan int) {
	c <- 1
	close(c)
}

func main() {
	ch := make(chan int)
	go addValue(ch)
	fmt.Println(<-ch)
}
