package main

import (
	"fmt"
)

func main() {
	var r, c, ans, sum, tmp int
	fmt.Scan(&r, &c)
	sens := make([][]int, r)
	for i := range sens {
		sens[i] = make([]int, c)
		for j := range sens[i] {
			fmt.Scan(&sens[i][j])
		}
	}
	for b := 0; b < (1 << uint(r)); b++ {
		tmp = 0
		for j := 0; j < c; j++ {
			sum = 0
			for i := 0; i < r; i++ {
				if (b&(1<<uint(i)) == 0 && sens[i][j] == 1) || (b&(1<<uint(i)) != 0 && sens[i][j] == 0) {
					sum++
				}
			}
			if sum > r/2 {
				tmp += sum
			} else {
				tmp += r - sum
			}
		}
		if tmp > ans {
			ans = tmp
		}
	}
	fmt.Println(ans)
}
