package main

import (
	"fmt"
)

type Z struct {
	x, y float32
}

func main() {
	var n, ans int
	fmt.Scan(&n)
	M := make(map[Z]int)
	var x, y float32
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		M[Z{x, y}] = 1
	}
	for z1, _ := range M {
		for z2, _ := range M {
			if z1 != z2 {
				m1 := (z1.x + z2.x) / 2
				m2 := (z1.y + z2.y) / 2
				ax := (z2.x - z1.x) / 2
				ay := (z2.y - z1.y) / 2
				z3 := Z{m1 + ay, m2 - ax}
				z4 := Z{m1 - ay, m2 + ax}
				if _, ok := M[z3]; ok {
					if _, ok := M[z4]; ok {
						s := (z1.x-z3.x)*(z1.x-z3.x) + (z1.y-z3.y)*(z1.y-z3.y)
						if int(s) > ans {
							ans = int(s)
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
