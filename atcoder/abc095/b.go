package main

import (
	"fmt"
)

func main() {
	var n, x, m int
	fmt.Scan(&n, &x)
	l := 1001
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		x -= m
		if l > m {
			l = m
		}
	}
	fmt.Println(x/l + n)
}
