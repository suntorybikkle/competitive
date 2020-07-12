package main

import (
	"fmt"
)

func main() {
	var H, W, K, ans int
	fmt.Scan(&H, &W, &K)
	mas := make([]string, H)
	for h := range mas {
		fmt.Scan(&mas[h])
	}
	for b := 0; b < (1 << (H + W)); b++ {
		k := 0
		for h := 0; h < H; h++ {
			if (b>>h)&1 == 1 {
				continue
			}
			for w := 0; w < W; w++ {
				if (b>>(H+w))&1 != 1 && string(mas[h][w]) == "#" {
					k += 1
				}
			}
		}
		if k == K {
			ans += 1
		}
	}
	fmt.Println(ans)
}
