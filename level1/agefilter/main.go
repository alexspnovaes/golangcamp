package main

import "fmt"

func ageFilter(minAge int, maxAge int, ages []int) []int {
	var filteredAges []int
	for _, age := range ages {
		if (age >= minAge) && (age <= maxAge) {
			filteredAges = append(filteredAges, age)
		}
	}
	return filteredAges
}

func main() {
	ages := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for _, age := range ageFilter(1, 10, ages) {
		fmt.Println(age)
	}
}
