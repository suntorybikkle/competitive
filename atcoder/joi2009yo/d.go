package main

import (
	"fmt"
)

var space [][]int
var done [][]bool

func dfs(x, y, l int) int {
	if space[y][x] == 0 || done[y][x] {
		return l
	}
	done[y][x] = true
	l++
	nl := l
	nl = max(nl, dfs(x, y-1, l))
	nl = max(nl, dfs(x, y+1, l))
	nl = max(nl, dfs(x-1, y, l))
	nl = max(nl, dfs(x+1, y, l))
	done[y][x] = false
	return nl
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	space = make([][]int, n+2)
	space[0], space[n+1] = make([]int, m+2), make([]int, m+2)
	done = make([][]bool, n+2)
	for i := range done {
		done[i] = make([]bool, m+2)
	}
	for i := range space[2:] {
		space[i+1] = make([]int, m+2)
		for j := range space[i+1][2:] {
			fmt.Scan(&space[i+1][j+1])
		}
	}
	ans := 0
	for i := range space[2:] {
		for j := range space[i][2:] {
			ans = max(ans, dfs(j, i, 0))
		}
	}
	fmt.Println(ans)
}
