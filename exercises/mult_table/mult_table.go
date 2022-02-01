package main

import "fmt"

func MultTable() (s [10][10]int) {
	for i := range s {
		for j := range s[i] {
			s[i][j] = i * j
		}
	}
	return s
}

func main() {
	table := MultTable()
	for i := range table {
		for j := range table[i] {
			fmt.Printf("%3d", table[i][j])
		}
		fmt.Println()
	}
}
