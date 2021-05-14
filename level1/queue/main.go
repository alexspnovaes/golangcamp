package main

import "fmt"

type q struct {
	Queue []int
}

func (q *q) addItem(item int) {
	q.Queue = append(q.Queue, item)
}

func (q *q) pop() {
	q.Queue = q.Queue[1:]
}

func (q *q) listItems() {
	fmt.Println("--------")
	for _, item := range q.Queue {
		fmt.Println(item)
	}
	fmt.Println("--------")
}

func main() {
	queue := new(q)
	queue.addItem(1)
	queue.addItem(2)
	queue.addItem(3)
	queue.addItem(4)
	queue.addItem(5)
	queue.listItems()

	queue.pop()
	queue.listItems()

	queue.pop()
	queue.listItems()
}
