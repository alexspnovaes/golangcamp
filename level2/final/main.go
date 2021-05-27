package main

import "fmt"

type General struct {
	name  string
	price float64
}

type Book struct {
	General
}

type Game struct {
	General
}

type Product interface {
	PrintInformation(i interface{})
	ApplyDiscount(i interface{})
}

func (g General) PrintInformation(i interface{}) {

	switch i.(type) {
	case Game:
		fmt.Println("Game")
	case Book:
		fmt.Println("Book")
	}

	fmt.Println(g.name)
	fmt.Println(g.price)
}

func (g General) ApplyDiscount(i interface{}) {
	switch i.(type) {
	case Game:
		fmt.Println("Applied discount of 20%")
	case Book:
		fmt.Println("Applied discount of 10%")
	}
}

func main() {
	a := Book{General{
		name:  "title of book",
		price: 10,
	}}

	a.PrintInformation(a)
	a.ApplyDiscount(a)

	b := Game{General{
		name:  "GTA 5",
		price: 60,
	}}

	b.PrintInformation(b)
	b.ApplyDiscount(b)
}
