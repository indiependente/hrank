package main

import (
	"fmt"
	"hrank/utils"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ratings := createSliceOfSize(n)
	fmt.Println(getMinCandies(ratings))
}

func getMinCandies(ratings []int) int {
	n := len(ratings)
	min, max := utils.MinMax(ratings)
	sToC := make(map[int][]int) // score to candies hashmap
	prevR := -1                 // previous rating
	prevS := -1                 // previous score
	for _, e := range ratings {
		candies := 0
		if len(sToC[e]) == 0 {

		}
		sToC[e] = append(sToC[e])
	}
	return sumArrayOfArray(getValuesFromMap(sToC))
}

func createSliceOfSize(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}

func getValuesFromMap(m map[int][]int) [][]int {
	var values [][]int
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func sumArrayOfArray(array [][]int) int {
	aSum := 0
	for _, a := range array {
		aSum += utils.Sum(a)
	}
	return aSum
}
