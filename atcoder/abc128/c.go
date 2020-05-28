package main

import (
	"fmt"
)

func main() {
	var n, m, k, c, ans int
	fmt.Scan(&n, &m)
	corrs := make([][]int, m)
	for i := range corrs {
		fmt.Scan(&k)
		corrs[i] = make([]int, k)
		for j := range corrs[i] {
			fmt.Scan(&corrs[i][j])
		}
	}
	p := make([]int, m)
	for i := range p {
		fmt.Scan(&p[i])
	}
	for b := 0; b < (1 << uint(n)); b++ {
		f := true
		for i, corr := range corrs {
			c = 0
			for _, s := range corr {
				if b&(1<<uint(s-1)) != 0 {
					c += 1
				}
			}
			if c%2 != p[i] {
				f = false
			}
		}
		if f {
			ans++
		}
	}
	fmt.Println(ans)
}
