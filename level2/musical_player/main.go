package main

import "fmt"

type Trumpeter struct {
	Name string
}

type Violinist struct {
	Name string
}

type MusicalPlayer interface {
	Greetings()
}

func (t Trumpeter) Greetings() {
	fmt.Println(t.Name)
}

func (t Violinist) Greetings() {
	fmt.Println(t.Name)
}

func main() {
	var jhon MusicalPlayer = Violinist{Name: "Jhon"}

	var elisa MusicalPlayer = Trumpeter{Name: "Elisa"}

	var mPlayers = []MusicalPlayer{jhon, elisa}

	for _, j := range mPlayers {
		j.Greetings()
	}
}
