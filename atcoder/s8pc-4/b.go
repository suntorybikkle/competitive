package main

import (
	"fmt"
)

func main() {
	var n, k, colors int
	var minH, tmp, ans int64
	ans = 1000000000 * 20
	fmt.Scan(&n, &k)
	A := make([]int64, n)
	for i := range A {
		fmt.Scan(&A[i])
	}
	for b := 0; b < (1 << uint(n)); b++ {
		colors = 0
		tmp = 0
		for i, a := range A {
			if i == 0 || a > minH {
				minH = a
				colors += 1
				continue
			}
			if b&(1<<uint(i)) != 0 {
				minH = minH + 1
				tmp += minH - a
				colors += 1
			}
		}
		if ans > tmp && k == colors {
			ans = tmp
		}
	}
	fmt.Println(ans)
}
