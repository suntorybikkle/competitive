package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	ans := 700
	for _, s := range s {
		if string(s) == "o" {
			ans += 100
		}
	}
	fmt.Println(ans)
}
