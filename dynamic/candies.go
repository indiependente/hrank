package main

import (
	"fmt"
	"hrank/utils"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	ratings := utils.CreateSliceOfSize(n)
	fmt.Println(getMinCandies(ratings))
}

func getMinCandies(ratings []int) int {
	lr := len(ratings)
	candies := make([]int, lr)

	// everyone gets a candy
	for i := 0; i < lr; i++ {
		candies[i] = 1
	}

	// first pass - increasing sequence
	for i := 1; i < lr; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// second pass - decreasing sequence
	for i := lr - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] < candies[i+1]+1 {
			candies[i] = candies[i+1] + 1
		}
	}
	return utils.Sum(candies)
}
