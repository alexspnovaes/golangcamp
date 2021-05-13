package main

import "fmt"

type Product struct {
	id   int
	name string
}

type Inventory struct {
	products map[int]Product
}

func addProduct(p Product, i *Inventory) {
	if p.id == 0 {
		fmt.Println("id must have value")
	} else if v, found := i.products[p.id]; found {
		fmt.Println("id already exists", v)
	} else {
		i.products[p.id] = p
	}
}

func main() {
	inventory := new(Inventory)
	inventory.products = make(map[int]Product)

	addProduct(Product{name: "Product1"}, inventory)
	addProduct(Product{id: 1, name: "Product2"}, inventory)
	addProduct(Product{id: 1, name: "Product3"}, inventory)
	addProduct(Product{id: 2, name: "Product4"}, inventory)
	addProduct(Product{id: 3, name: "Product5"}, inventory)

	for _, p := range inventory.products {
		fmt.Println("Id:", p.id, "=>", "Name:", p.name)
	}
}
