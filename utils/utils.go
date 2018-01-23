package utils

import "fmt"

// CreateSliceOfSize returns a slice of ints which elements are taken in input from Stdin
func CreateSliceOfSize(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}

// CreateInt64SliceOfSize returns a slice of int64 which elements are taken in input from Stdin
func CreateInt64SliceOfSize(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}

// Sum returns the sum of the integers in the input array
func Sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// SumInt64 returns the sum of the int64 elements in the input array
func SumInt64(arr []int64) int64 {
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// SelectionSortInt64 in-place selection sort for int64 elements
func SelectionSortInt64(array []int64) []int64 {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
