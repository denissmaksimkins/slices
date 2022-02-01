package main

import (
	"fmt"
)

func FilterEvens(s []int) (d []int) {
	for _, v := range s {
		if v%2 == 0 {
			d = append(d, v)
		}
	}
	return d
}

func FilterOdds(s []int) (a []int) {
	for _, v := range s {
		if v%2 != 0 {
			a = append(a, v)

		}

	}
	return a
}

func main() {
	fmt.Println("Enter slice size:")
	var size uint
	fmt.Scan(&size)
	input := make([]int, size)
	for i := range input {
		fmt.Scan(&input[i])
	}
	fmt.Println(input)
	fmt.Println(FilterEvens(input), "Odd numbers")
	fmt.Println(FilterOdds(input), "Even numbers")
}
