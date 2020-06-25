package main

import (
	"fmt"
)

func main() {
	var h, w, c, ans int
	fmt.Scan(&h, &w)
	grid := make([][]int, h)
	var S string
	for j := range grid {
		fmt.Scan(&S)
		grid[j] = make([]int, w)
		for i, s := range S {
			if string(s) == "#" {
				grid[j][i] = -1
			} else {
				c++
				grid[j][i] = 0
			}
		}
	}
	grid[0][0] = 1
	q := make([][]int, 0)
	q = append(q, []int{0, 0})
	for len(q) != 0 {
		y, x := q[0][0], q[0][1]
		q = q[1:]
		for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			ny, nx := y+d[0], x+d[1]
			if ny < 0 || ny == h || nx < 0 || nx == w {
				continue
			}
			if grid[ny][nx] == 0 {
				grid[ny][nx] = grid[y][x] + 1
				q = append(q, []int{ny, nx})
			}
		}
	}
	ans = c - grid[h-1][w-1]
	if grid[h-1][w-1] == 0 {
		ans = -1
	}
	fmt.Println(ans)
}
