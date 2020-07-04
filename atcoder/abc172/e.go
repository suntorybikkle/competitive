package main

import (
	"fmt"
)

const MOD = 1e9 + 7

var fact [5e5 + 1]int
var finv [5e5 + 1]int
var inv [5e5 + 1]int

func comb(n, k int) int {
	return fact[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

func perm(n, k int) int {
	return fact[n] * (finv[n-k] % MOD) % MOD
}

func main() {
	var n, m, ans int
	fmt.Scan(&n, &m)
	fact[0], fact[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < m+1; i++ {
		fact[i] = i * fact[i-1] % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
	s := -1
	for k := 0; k <= n; k++ {
		s *= -1
		ans += comb(n, k) * perm(m, k) % MOD * perm(m-k, n-k) % MOD * perm(m-k, n-k) % MOD * s
		ans %= MOD
		if ans < 0 {
			ans += MOD
		}
	}
	fmt.Println(ans)
}
