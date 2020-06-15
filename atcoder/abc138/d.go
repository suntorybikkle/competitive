package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var G map[int][]int
var X []int
var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func dfs(v, p int) {
	for _, r := range G[v] {
		if r == p {
			continue
		}
		X[r-1] += X[v-1]
		dfs(r, v)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	q := nextInt()
	G = make(map[int][]int)
	for i := 1; i < n; i++ {
		a, b := nextInt(), nextInt()
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	X = make([]int, n)
	for i := 0; i < q; i++ {
		X[nextInt()-1] += nextInt()
	}
	dfs(1, -1)
	for _, x := range X {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}
