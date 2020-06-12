package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, ans int
	fmt.Scan(&n, &m)
	p := make([]int, n+1)
	for i := range p[1:] {
		fmt.Scan(&p[i])
	}
	var twoP []int
	twoSet := make(map[int]bool)
	for i := range p {
		for j := range p {
			two := p[i] + p[j]
			if _, ok := twoSet[two]; ok {
				continue
			}
			twoSet[two] = true
			twoP = append(twoP, two)
		}
	}
	sort.Ints(twoP)
	for _, two := range twoP {
		if two > m {
			continue
		}
		left, right := 0, len(twoP)-1
		for right-1 > left {
			mean := (right + left) / 2
			if two+twoP[mean] > m {
				right = mean
			} else {
				left = mean
			}
		}
		if twoP[left]+two > ans {
			ans = twoP[left] + two
		}
	}
	fmt.Println(ans)
}
