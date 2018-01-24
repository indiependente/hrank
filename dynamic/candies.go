package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ratings := createSliceOfSize(n)
	fmt.Println(getMinCandies(ratings))
}

func getMinCandies(ratings []int) ([]int, int) {
	lr := len(ratings)
	scores := make([]int, lr)
	var decr bool
	var dl, el int
	for i := 0; i < lr; i++ {
		if i != lr-1 && ratings[i] > ratings[i+1] {
			if !decr {
				decr = true
				dl = decrLen(ratings, i)
			}
			for j := 0; j < dl; j++ {
				scores[j] = dl - j + 1
			}
		} else if ratings[i] < ratings[i+1] && i != 0 {
			decr = false
			scores[i] = scores[i-1] + 1
		} else if ratings[i] == ratings[i+1] {
			decr = false
			el = equalLen(ratings, i)

		}
	}

	return scores, sum(scores)
}

func createSliceOfSize(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}

func sum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func decrLen(arr []int, pos int) int {
	length := 0
	for i := pos; i < len(arr)-1; i++ {
		if i != len(arr) && arr[i] < arr[i+1] {
			length++
		}
	}
	return length
}

func equalLen(arr []int, pos int) int {
	length := 0
	for i := pos; i < len(arr)-1; i++ {
		if i != len(arr) && arr[i] == arr[i+1] {
			length++
		}
	}
	return length
}
