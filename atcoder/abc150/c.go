package main

import (
	"fmt"
	"reflect"
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
	var n, a, b int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := range p {
		fmt.Scan(&p[i])
	}
	q := make([]int, n)
	for i := range q {
		fmt.Scan(&q[i])
	}
	order := make([]int, n)
	for i := range order {
		order[i] = i + 1
	}
	for c := 0; ; c++ {
		if reflect.DeepEqual(order, p) {
			a = c
		}
		if reflect.DeepEqual(order, q) {
			b = c
		}
		if !NextPermutation(sort.IntSlice(order)) {
			break
		}
	}
	fmt.Println(abs(a - b))
}
