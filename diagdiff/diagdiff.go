package main

import (
	"fmt"
	"hrank/utils"
	"log"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	m := readMatrix(n)
	d1, d2 := primaryDiag(m), secondaryDiag(m)
	fmt.Println(abs(utils.Sum(d1) - utils.Sum(d2)))
}

func primaryDiag(m [][]int) []int {
	n := len(m)
	diag := make([]int, n)
	for i := 0; i < n; i++ {
		diag[i] = m[i][i]
	}
	return diag
}

func secondaryDiag(m [][]int) []int {
	n := len(m)
	diag := make([]int, n)
	for i := 0; i < n; i++ {
		diag[i] = m[i][n-i-1]
	}
	return diag
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func readMatrix(n int) [][]int {
	if n < 1 {
		log.Fatal("n must be at least 1.")
	}
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		m[i] = utils.CreateSliceOfSize(n)
		for j := 0; j < n; j++ {
			if m[i][j] < -100 || m[i][j] > 100 {
				log.Fatal("m[i][j] must be between -100 and 100.")
			}
		}
	}
	return m
}
