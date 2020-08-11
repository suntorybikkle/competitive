package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	var n, k int
	fmt.Scan(&n, &k)
	A := make([]int, n)
	for i := range A {
		A[i] = nextInt()
	}
	lb, ub := 0, int(1e9+1)
	for ub-lb > 1 {
		mid := (lb + ub) / 2
		c := 0
		for _, a := range A {
			tmp := a / mid
			if a%mid == 0 {
				tmp--
			}
			c += tmp
		}
		if k >= c {
			ub = mid
		} else {
			lb = mid
		}
	}
	fmt.Println(ub)
}
