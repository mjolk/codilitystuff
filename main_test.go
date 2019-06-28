package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	test := []int{2, 3, 4, 8, 1, 5}

	sum := solution(test, 3, 2)
	fmt.Printf("found sum %d\n", sum)
	if sum != 21 {
		t.Fail()
	}

	test2 := []int{1, 5, 25, 8, 90, 21}
	sum = solution(test2, 2, 3)
	fmt.Printf("found sum %d\n", sum)

	sum = solution(test2, 3, 2)
	fmt.Printf("found sum %d\n", sum)
}
