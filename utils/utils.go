package utils

import "fmt"

// Note the use of << to create an untyped constant.
const bitsPerWord = 32 << uint(^uint(0)>>63)

//BitsPerWord is an implementation-specific size of int and uint in bits.
const BitsPerWord = bitsPerWord // either 32 or 64

// Implementation-specific integer limit values.
const (
	MaxInt  = 1<<(BitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1            // either -1 << 31 or -1 << 63
	MaxUint = 1<<BitsPerWord - 1     // either 1<<32 - 1 or 1<<64 - 1
)

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

// SelectionSort in-place selection sort for int64 elements
func SelectionSort(array []int) []int {
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

// Min returns the min element of the input slice and the position where it has been found
func Min(a []int) (int, int) {
	min := MaxInt
	pos := -1
	for i, e := range a {
		if e < min {
			min = e
			pos = i
		}
	}
	return min, pos
}

// Max returns the max element of the input slice and the position where it has been found
func Max(a []int) (int, int) {
	max := MinInt
	pos := -1
	for i, e := range a {
		if e > max {
			max = e
			pos = i
		}
	}
	return max, pos
}

// MinMax returns the minimum and the maximum value of the array in input
func MinMax(a []int) (int, int) {
	max := MinInt
	min := MaxInt
	for _, e := range a {
		if e > max {
			max = e
		}
		if e < min {
			min = e
		}
	}
	return min, max
}
