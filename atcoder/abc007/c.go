package main

import (
	"fmt"
)

func main() {
	var r, c, sy, sx, gy, gx int
	fmt.Scan(&r, &c, &sy, &sx, &gy, &gx)
	maze := make([][]int, r)
	var S string
	for i := range maze {
		fmt.Scan(&S)
		maze[i] = make([]int, c)
		for j, s := range S {
			if string(s) == "#" {
				maze[i][j] = -2
			} else {
				maze[i][j] = -1
			}
		}
	}
	maze[sy-1][sx-1] = 0
	q := make([][]int, 0)
	q = append(q, []int{sy - 1, sx - 1})
	for len(q) != 0 {
		y, x := q[0][0], q[0][1]
		q = q[1:]
		for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
			ny, nx := y+d[0], x+d[1]
			if maze[ny][nx] == -1 {
				maze[ny][nx] = maze[y][x] + 1
				q = append(q, []int{ny, nx})
			}
		}
	}
	fmt.Println(maze[gy-1][gx-1])
}
