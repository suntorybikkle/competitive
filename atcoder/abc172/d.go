package main

import (
	"fmt"
)

func main() {
	var n, ans int64
	fmt.Scan(&n)
	F := make([]int64, n+1)
	for i := int64(1); i < n+1; i++ {
		for j := i; j < n+1; j += i {
			F[j] += j
		}
	}
	for _, f := range F {
		ans += f
	}
	fmt.Println(ans)
}
