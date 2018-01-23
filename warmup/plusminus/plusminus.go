package main

import (
	"fmt"
	"hrank/utils"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := utils.CreateSliceOfSize(n)
	pos, neg, zero := posNegZeroFrac(s)
	fmt.Println(pos)
	fmt.Println(neg)
	fmt.Println(zero)
}

func posNegZeroFrac(arr []int) (float32, float32, float32) {
	p, n, z := 0, 0, 0
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] > 0 {
			p++
		} else if arr[i] < 0 {
			n++
		} else {
			z++
		}
	}
	fLength := float32(length)
	return float32(p) / fLength, float32(n) / fLength, float32(z) / fLength
}
