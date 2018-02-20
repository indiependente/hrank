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
	if (v2 != v1) && ((x1-x2)/(v2-v1)) > 0 {
		fmt.Print("YES")
	} else if (v2 == v1) && (x1-x2) > 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
