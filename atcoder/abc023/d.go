package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func judge(hs [][]int, x int64) bool {
	o := make([]int, len(hs))
	for i, _ := range o {
		if int64(hs[i][0]) > x {
			return false
		}
		o[i] = int((x - int64(hs[i][0])) / int64(hs[i][1]))
	}
	sort.Ints(o)
	for i, _ := range o {
		if i > o[i] {
			return false
		}
	}
	return true
}

var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()

	hs := make([][]int, n)
	for i := range hs {
		hs[i] = make([]int, 2)
		hs[i][0] = readInt()
		hs[i][1] = readInt()
	}
	left, right := int64(0), int64(1000000000*100003)
	var mean int64
	for right-1 > left {
		mean = (right + left) / 2
		if judge(hs, mean) {
			right = mean
		} else {
			left = mean
		}
	}
	fmt.Println(right)
}
