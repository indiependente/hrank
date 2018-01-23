package main

import (
	"fmt"
	"hrank/utils"
	"os"
)

func main() {
	n := readN()
	_, max, occurr := createSliceOfSizeAndGetMax(n)
	fmt.Println(occurr[max])
}

func readN() int {
	var n int
	fmt.Scanf("%d", &n)
	if n < 1 || n > 100000 {
		os.Exit(1)
	}
	return n
}

func createSliceOfSizeAndGetMax(n int) ([]int, int, map[int]int) {
	a := make([]int, n)
	m := make(map[int]int)
	max := utils.MinInt
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		m[a[i]]++
		if a[i] > max {
			max = a[i]
		}
	}
	return a, max, m
}
