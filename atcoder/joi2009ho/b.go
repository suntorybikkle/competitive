package main

import (
	"fmt"
	"sort"
)

func bisec(S []int, k int) int {
	lIdx, rIdx := 0, len(S)-1
	mIdx := rIdx / 2
	for {
		if S[mIdx+1] > k && k >= S[mIdx] {
			break
		} else if S[mIdx] > k {
			rIdx = mIdx
			mIdx = (rIdx + lIdx) / 2
		} else {
			lIdx = mIdx
			mIdx = (rIdx + lIdx) / 2
		}
	}
	l, r := k-S[mIdx], S[mIdx+1]-k
	if r > l {
		return l
	} else {
		return r
	}
}

func main() {
	var d, n, m int
	var ans int64
	fmt.Scan(&d, &n, &m)
	S := make([]int, n)
	for i := range S[1:] {
		fmt.Scan(&S[i+1])
	}
	S = append(S, d)
	K := make([]int, m)
	for i := range K {
		fmt.Scan(&K[i])
	}
	sort.Ints(S)
	for _, k := range K {
		ans += int64(bisec(S, k))
	}
	fmt.Println(ans)
}
