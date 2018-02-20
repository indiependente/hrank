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
	var t float32
	fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)

	t = float32(x1-x2) / float32(v2-v1)

	if t >= 0 && t == float32(int(t)) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
