package main

import (
	"fmt"
	"hrank/utils"
)

func main() {
	array := utils.CreateInt64SliceOfSize(5)
	_ = utils.SelectionSortInt64(array)
	fmt.Printf("%d %d\n", utils.SumInt64(array[:len(array)-1]), utils.SumInt64(array[1:]))
}
