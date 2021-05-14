package main

import "fmt"

type Stack struct {
	items []int
}

func (q *Stack) addItem(item int) {
	q.items = append(q.items, item)
}

func (q *Stack) pop() {
	q.items = q.items[1:]
}

func (q *Stack) peek() {
	fmt.Println(q.items[0])
}

func (q *Stack) listItems() {
	fmt.Println("--------")
	for _, item := range q.items {
		fmt.Println(item)
	}
	fmt.Println("--------")
}

func main() {
	items := new(Stack)
	items.addItem(1)
	items.addItem(2)
	items.addItem(3)
	items.listItems()

	items.pop()
	items.listItems()

	items.addItem(4)
	items.listItems()

	items.peek()
}
