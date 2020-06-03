package main

import (
	"fmt"
)

func fact(x int) int {
	res := 1
	for x > 0 {
		res *= x
		x--
	}
	return res
}

func dictorder(array []int) int {
	res, length := 1, len(array)
	done := make([]int, length+1)
	for i, x := range array {
		t := 0
		for y := 1; y < x; y++ {
			if done[y] == 0 {
				t++
			}
		}
		res += fact(length-i-1) * t
		done[x] = 1
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}
	q := make([]int, n)
	for i := range q {
		fmt.Scan(&q[i])
	}
	fmt.Println(abs(dictorder(p) - dictorder(q)))
}
