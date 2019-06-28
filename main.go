package main

import "fmt"

func main() {
	fmt.Println("codility stuff")
}

func solution(src []int, k, l int) int {
	if k+l > len(src) {
		return -1
	}

	sum1, kstart := maxInterval(src, k)
	src2 := append(src[:kstart], src[kstart+k:]...)
	fmt.Printf("src2 => %+v\n", src2)
	sum2, _ := maxInterval(src2, l)
	return (sum1 + sum2) % 1000000007

}

func maxInterval(src []int, k int) (int, int) {

	var fsum int
	for i := 0; i < k; i++ {
		fsum += src[i]
	}

	sum := fsum
	var kstart int
	for j := k; j < len(src); j++ {
		sum += src[j] - src[j-k]
		if sum > fsum {
			kstart++
			fsum = sum
		}
	}
	fmt.Printf("sum %d, kstart: %d\n", fsum, kstart)
	return fsum, kstart
}
