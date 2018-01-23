package main

import (
	"fmt"
	"hrank/utils"
	"log"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n < 1 || n > 10 {
		log.Fatal("N must be between 1 and 10.")
	}
	array := utils.CreateInt64SliceOfSize(n)
	if !isValid(array) {
		log.Fatal("Values between 0 and 10^10.")
	}
	fmt.Println(sum(array))
}

func sum(arr []int64) int64 {
	var s int64
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

func isValid(arr []int64) bool {
	maxAllowed := int64(10000000000)
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 || arr[i] > maxAllowed {
			return false
		}
	}
	return true
}
