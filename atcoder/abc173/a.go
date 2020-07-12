package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	for {
		if 1000 >= n {
			break
		}
		n -= 1000
	}
	fmt.Println(1000 - n)
}
