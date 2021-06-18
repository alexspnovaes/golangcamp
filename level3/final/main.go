package main

import "fmt"

func filterByAgeRange(fromAge, toAge int, ages []int) []int {
	response := make([]int, 0)
	for _, ageToCheck := range ages {
		if fromAge <= ageToCheck && ageToCheck <= toAge {
			response = append(response, ageToCheck)
		}
	}
	return response
}

func main() {
	ages := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	filteedages := filterByAgeRange(2, 15, ages)

	for _, age := range filteedages {
		fmt.Println(age)
	}
}
