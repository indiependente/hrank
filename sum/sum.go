package main

import (
	"fmt"
	"hrank/utils"
)

func main() {
	var n int
	fmt.Scan(&n)
	array := utils.CreateSliceOfSize(n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += array[i]
	}
	fmt.Println(sum)
}
