package main

import (
	"fmt"
)

func min(n, m int) int {
	if n >= m {
		return m
	} else {
		return n
	}
}

func main() {
	var a, b, c, x, y int
	fmt.Scan(&a, &b, &c, &x, &y)
	m := min(a+b, 2*c)
	if x >= y {
		fmt.Println(y*m + (x-y)*min(a, 2*c))
	} else {
		fmt.Println(x*m + (y-x)*min(b, 2*c))
	}
}
