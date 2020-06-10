package main

import (
	"fmt"
	"math"
)

func main() {
	var p float64
	fmt.Scan(&p)
	x := 1.5 * math.Log2(p*math.Log(2)/1.5)
	if x > 0 {
		fmt.Println(x + 1.5/math.Log(2))
	} else {
		fmt.Println(p)
	}
}
