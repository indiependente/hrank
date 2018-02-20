/*
x1 + i*v1 = x2 + i*v2
x1 - x2 = i*v2 - i*v1
x1 - x2 = i*(v2 - v1)
i = (x1 - x2)/(v2 - v1)
*/

package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)
	i := (x1 - x2) / (v2 - v1)
	fmt.Println(x1 + i*v1)
	if ((x1 - x2) / (v2 - v1)) >= 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
