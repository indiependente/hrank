package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Printf("%s", " ")
		}
		for j := 0; j < i+1; j++ {
			fmt.Printf("%s", "#")
		}
		fmt.Printf("\n")

	}
}
