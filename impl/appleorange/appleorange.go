package main

import "fmt"

type problem struct {
	s       int
	t       int
	a       int
	b       int
	m       int
	n       int
	apples  []int
	oranges []int
}

func main() {
	p := readProblem()
	// fmt.Printf("%v\n", p)
	s := solveProblem(p)

	fmt.Printf("%d\n%d\n", s[0], s[1])
}

func solveProblem(p *problem) [2]int {
	result := [2]int{}
	for i := 0; i < p.m; i++ {
		if p.a+p.apples[i] >= p.s && p.a+p.apples[i] <= p.t {
			result[0]++
		}
	}
	for i := 0; i < p.n; i++ {
		if p.b+p.oranges[i] >= p.s && p.b+p.oranges[i] <= p.t {
			result[1]++
		}
	}
	return result
}

func readProblem() *problem {
	p := &problem{}
	fmt.Scanf("%d %d", &p.s, &p.t)
	fmt.Scanf("%d %d", &p.a, &p.b)
	fmt.Scanf("%d %d", &p.m, &p.n)
	p.apples = make([]int, p.m)
	p.oranges = make([]int, p.n)
	for i := 0; i < p.m; i++ {
		fmt.Scanf("%d", &p.apples[i])
	}
	for i := 0; i < p.n; i++ {
		fmt.Scanf("%d", &p.oranges[i])
	}
	return p
}
