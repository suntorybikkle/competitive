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
	A := make([]int, n)
	for i := range A {
		A[i] = nextInt()
	}
	sort.Ints(A)
	S := make([][]bool, 1000001)
	for i := range S {
		S[i] = make([]bool, 2)
	}
	for _, a := range A {
		if !S[a][1] {
			if S[a][0] {
				S[a][0] = false
				S[a][1] = true
				continue
			}
			S[a][0] = true
			for j := 2; 1000001 > j*a; j++ {
				S[j*a][1] = true
			}
		}
	}
	ans := 0
	for _, a := range A {
		if S[a][0] {
			ans += 1
		}
	}
	fmt.Println(ans)
}
