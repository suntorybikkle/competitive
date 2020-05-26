package main

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if 0 > x {
		return -x
	}
	return x
}

func f(mas []int) int64 {
	// sort.Slice(mas, func(i, j int) bool { return mas[i] < mas[j] })
	sort.Ints(mas)
	mean := mas[len(mas)/2]
	var sum int64
	for _, m := range mas {
		sum += int64(abs(m - mean))
	}
	return sum
}

func main() {
	var n int
	var ans int64
	fmt.Scan(&n)
	A := make([]int, n)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i], &B[i])
		ans += int64(B[i] - A[i])
	}
	ans += f(A) + f(B)
	fmt.Println(ans)
}
