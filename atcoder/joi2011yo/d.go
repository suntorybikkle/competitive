package main

import (
	"fmt"
)

var S []int64
var n int
var dp [][]int64

func main() {
	fmt.Scan(&n)
	S = make([]int64, n)
	for i := range S {
		fmt.Scan(&S[i])
	}
	dp = make([][]int64, n)
	for i := range dp {
		dp[i] = make([]int64, 21)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	fmt.Println(dfs(1, S[0]))
}

func dfs(i int, sum int64) int64 {
	if dp[i][sum] != -1 {
		return dp[i][sum]
	}
	if i == n-1 {
		if S[i] == sum {
			return 1
		}
		return 0
	}
	var tmp int64
	tmp = 0
	if sum+S[i] <= 20 {
		tmp += dfs(i+1, sum+S[i])
	}
	if sum-S[i] >= 0 {
		tmp += dfs(i+1, sum-S[i])
	}
	dp[i][sum] = tmp
	return tmp
}
