package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n := nextInt()
	m := nextInt()
	k := nextInt()
	A := make([]int, n+1)
	B := make([]int, m+1)
	for i := range A[1:] {
		A[i+1] = nextInt() + A[i]
	}
	for i := range B[1:] {
		B[i+1] = nextInt() + B[i]
	}
	var ans int
	for i, a := range A {
		if a > k {
			continue
		}
		tmp := i + sort.SearchInts(B, k-a+1) - 1
		if tmp > ans {
			ans = tmp
		}
	}
	fmt.Println(ans)
}
