package main

import (
	"fmt"
)

func main() {
	var w, h, ans int
	fmt.Scan(&w, &h)
	space := make([][]int, h)
	for j := range space {
		space[j] = make([]int, w)
		for i := range space[j] {
			fmt.Scan(&space[j][i])
		}
	}
	dy := []int{0, 0, -1, 1, -1, 1}
	for j := range space {
		for i := range space[j] {
			if space[j][i] != 1 {
				continue
			}
			q := make([][]int, 0)
			q = append(q, []int{j, i})
			space[j][i] = 2
			for len(q) != 0 {
				y, x := q[0][0], q[0][1]
				q = q[1:]
				var dx []int
				if y%2 == 0 {
					dx = []int{-1, 1, 0, 0, 1, 1}
				} else {
					dx = []int{-1, 1, -1, -1, 0, 0}
				}
				c := 0
				for d := range dy {
					ny, nx := y+dy[d], x+dx[d]
					if ny == h || ny < 0 || nx == w || nx < 0 {
						continue
					}
					if space[ny][nx] == 0 {
						continue
					}
					c++
					if space[ny][nx] == 1 {
						space[ny][nx] = 2
						q = append(q, []int{ny, nx})
					}
				}
				ans += 6 - c
			}
		}
	}
	for j := range space {
		for i := range space[j] {
			if space[j][i] != 0 {
				continue
			}
			q := make([][]int, 0)
			q = append(q, []int{j, i})
			space[j][i] = 1
			var wall bool
			rc := 0
			for len(q) != 0 {
				y, x := q[0][0], q[0][1]
				q = q[1:]
				var dx []int
				if y%2 == 0 {
					dx = []int{-1, 1, 0, 0, 1, 1}
				} else {
					dx = []int{-1, 1, -1, -1, 0, 0}
				}
				c := 0
				for d := range dy {
					ny, nx := y+dy[d], x+dx[d]
					if ny == h || ny < 0 || nx == w || nx < 0 {
						wall = true
						continue
					}
					if space[ny][nx] == 2 {
						continue
					}
					c++
					if space[ny][nx] == 0 {
						space[ny][nx] = 1
						q = append(q, []int{ny, nx})
					}
				}
				rc += 6 - c
			}
			if !wall {
				ans -= rc
			}
		}
	}
	fmt.Println(ans)
}
