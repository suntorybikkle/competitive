package main

import (
	"fmt"
)

func main() {
	var x int
	for i := 1; i < 6; i++ {
		fmt.Scan(&x)
		if x == 0 {
			fmt.Println(i)
		}
	}
}
