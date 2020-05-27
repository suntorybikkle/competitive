package main

import (
	"fmt"
)

func main() {
	var m, n, x, y int
	fmt.Scan(&m)
	fmt.Scan(&x, &y)
	org := coord{x, y}
	sign := make([]coord, m-1)
	for i := range sign {
		fmt.Scan(&x, &y)
		sign[i].x = x - org.x
		sign[i].y = y - org.y
	}
	fmt.Scan(&n)
	stars := make(map[coord]bool, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		stars[coord{x, y}] = true
	}
	for key, _ := range stars {
		contains := true
		for _, s := range sign {
			star := coord{key.x + s.x, key.y + s.y}
			if _, ok := stars[star]; !ok {
				contains = false
				break
			}
		}
		if contains {
			fmt.Printf("%d %d\n", key.x-org.x, key.y-org.y)
			break
		}
	}
}

type coord struct {
	x, y int
}
