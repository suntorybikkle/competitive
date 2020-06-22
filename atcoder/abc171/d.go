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
	A := make(map[int]int)
	ans, n := 0, nextInt()
	for i := 0; i < n; i++ {
		a := nextInt()
		ans += a
		A[a] += 1
	}
	q := nextInt()
	for i := 0; i < q; i++ {
		b := nextInt()
		c := nextInt()
		ans -= A[b] * b
		ans += A[b] * c
		A[c] += A[b]
		A[b] = 0
		fmt.Println(ans)
	}
}
