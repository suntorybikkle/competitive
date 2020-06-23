package main

import (
	"fmt"
	"strconv"
)

func main() {
	var h, w, n, ans int
	fmt.Scan(&h, &w, &n)
	city := make([]string, h)
	cheese := make([][]int, n+1)
	for j := range city {
		fmt.Scan(&city[j])
		for i, s := range city[j] {
			as := string(s)
			if as == "X" || as == "." {
				continue
			}
			if as == "S" {
				cheese[0] = []int{j, i}
			} else {
				is, _ := strconv.Atoi(as)
				cheese[is] = []int{j, i}
			}
		}
	}
	for i := 0; i < n; i++ {
		bfs := make([][]int, h)
		for b := range bfs {
			bfs[b] = make([]int, w)
		}
		q := make([][]int, 0)
		q = append(q, []int{cheese[i][0], cheese[i][1]})
		for len(q) != 0 {
			y, x := q[0][0], q[0][1]
			q = q[1:]
			for _, d := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				ny, nx := y+d[0], x+d[1]
				if ny == h || ny < 0 || nx == w || nx < 0 {
					continue
				}
				if string(city[ny][nx]) != "X" && bfs[ny][nx] == 0 {
					bfs[ny][nx] = bfs[y][x] + 1
					q = append(q, []int{ny, nx})
				}
			}
		}
		ans += bfs[cheese[i+1][0]][cheese[i+1][1]]
	}
	fmt.Println(ans)
}
