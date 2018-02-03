package main

import (
	"fmt"
)

func main() {
	var q, n int
	var m int64
	fmt.Scanf("%d", &q)
	fmt.Scanf("%d %d", &n, &m)
	array := createInt64SliceOfSize(n)
	sum := maxSubSum(array, n, m)
	fmt.Println(sum)
}

func maxSubSum(a []int64, n int, M int64) int64 {
	pref := make([]int64, n)
	var curr, max int64
	for i := 0; i < n; i++ {
		pref[i] = (curr + a[i]) % M
		curr = pref[i]
		if pref[i] > max {
			max = pref[i]
		}
	}
	return max
}

func createInt64SliceOfSize(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}
