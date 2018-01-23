package main

import (
	"fmt"
)

func main() {
	var hh, mm, ss int
	var t12 string
	fmt.Scanf("%d:%d:%d%s", &hh, &mm, &ss, &t12)
	if t12 == "PM" && hh != 12 {
		hh += 12
	} else if t12 == "AM" && hh == 12 {
		hh = 0
	}
	fmt.Printf("%02d:%02d:%02d\n", hh, mm, ss)
}
