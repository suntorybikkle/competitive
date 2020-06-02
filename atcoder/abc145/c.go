package main

import (
	"fmt"
	"math"
	"sort"
)

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	var n, sumx, sumy int
	var ans, c, tmp float64
	fmt.Scan(&n)

	coords := make([][]int, n)
	for i := range coords {
		coord := make([]int, 2)
		fmt.Scan(&coord[0], &coord[1])
		coords[i] = coord
	}
	order := make([]int, n)
	for i := range order {
		order[i] = i
	}
	for ; ; c++ {
		tmp = 0
		for i := 0; i < n-1; i++ {
			sumx = abs(coords[order[i]][0] - coords[order[i+1]][0])
			sumy = abs(coords[order[i]][1] - coords[order[i+1]][1])
			tmp += math.Sqrt(float64(sumx*sumx + sumy*sumy))
		}
		ans += tmp
		if !NextPermutation(sort.IntSlice(order)) {
			break
		}
	}
	fmt.Println(ans / (c + 1))
}
