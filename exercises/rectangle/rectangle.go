package main

import "fmt"

func Rect(x, y int) [][]int {
	rows, cols := x, y
	s := make([][]int, rows)
	for row := 0; row < rows; row++ {
		s[row] = make([]int, cols)
	}

	for i := 0; i < x; i++ {
		s[0][i] = 1
		s[y-1][i] = 1
	}
	for i := 0; i < y; i++ {
		s[i][0] = 1
		s[i][x-1] = 1
	}
	return s
}
func main() {
	var x, y int
	fmt.Println("Enter width and height:")
	fmt.Scan(&x, &y)
	if x <= 0 || y <= 0 {
		fmt.Println("Rectangle does not exist.")
		return
	}
	table := Rect(x, y)
	for i := range table {
		for j := range table[i] {
			fmt.Printf("%3d", table[i][j])
		}
		fmt.Println()
	}
}
