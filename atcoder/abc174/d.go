package main

import (
	"fmt"
)

func main() {
	var N, r, ans int
	var C string
	fmt.Scan(&N, &C)
	W := make([]int, 0)
	for i := range C {
		if string(C[N-1-i]) == "R" {
			r++
		} else {
			W = append(W, r)
		}
	}
	for i := range W {
		if W[len(W)-1-i] > ans {
			ans++
		}
	}
	fmt.Println(ans)
}
