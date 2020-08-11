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
	var n, d, ans int
	fmt.Scan(&n, &d)
	for i := 0; i < n; i++ {
		x, y := nextInt(), nextInt()
		if d*d >= x*x+y*y {
			ans++
		}
	}
	fmt.Println(ans)
}
