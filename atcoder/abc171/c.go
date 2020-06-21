package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)
	var S []int64
	for n > 0 {
		n -= 1
		S = append(S, n%26)
		n /= 26
	}
	for i := range S {
		fmt.Print(string("a"[0] + byte(S[len(S)-i-1])))
	}
	fmt.Println()
}
