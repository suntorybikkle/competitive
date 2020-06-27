package main

import (
	"fmt"
)

func main() {
	var ans int
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)
	for i, s := range S {
		if string(s) != string(T[i]) {
			ans++
		}
	}
	fmt.Println(ans)
}
