package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)
	n := 0
	for i := 1; i <= k; i++ {
		n = (n*10 + 7) % k
		if n == 0 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
