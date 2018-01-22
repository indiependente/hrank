package main

import (
	"fmt"
	"hrank/utils"
	"log"
)

type triplet [3]int

type player struct {
	t     triplet
	score int
}

func main() {
	t, err := readTriplet()
	if err != nil {
		log.Fatalf("Error in creating triplet: %s", err)
	}
	alice := player{t, 0}
	t, err = readTriplet()
	if err != nil {
		log.Fatalf("Error in creating triplet: %s", err)
	}
	bob := player{t, 0}
	comparePlayers(&alice, &bob)
	fmt.Printf("%d %d", alice.score, bob.score)
}

func readTriplet() (triplet, error) {
	s := utils.CreateSliceOfSize(3)
	if len(s) != 3 {
		return triplet{}, fmt.Errorf("Length error. Triplets must be of 3 values")
	}
	for _, sI := range s {
		if sI < 1 || sI > 100 {
			return triplet{}, fmt.Errorf("si violated constraint. Must be between 1 and 100.: si = %d", sI)
		}
	}
	t := sliceToTriplet(s)
	return t, nil
}

func comparePlayers(a, b *player) {
	for i := 0; i < 3; i++ {
		if a.t[i] < b.t[i] {
			b.score++
		} else if a.t[i] > b.t[i] {
			a.score++
		}
	}
}

func sliceToTriplet(s []int) triplet {
	t := triplet{}
	for i := 0; i < 3; i++ {
		t[i] = s[i]
	}
	return t
}
