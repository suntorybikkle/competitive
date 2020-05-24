package main

import (
	"fmt"
	"strconv"
)

func bi(m []int, l, r int) int {
	var ml, mr, mm int
	ml, mr = 0, len(m)-1
	mm = (ml + mr) / 2
	for {
		if m[mm] > l && r > m[mm] {
			return 1
		}
		if ml == mr {
			return 0
		} else if l >= m[mm] {
			ml = mm + 1
			mm = (ml + mr) / 2
		} else if m[mm] >= r {
			mr = mm
			mm = (ml + mr) / 2
		}
	}
}

func main() {
	var S string
	fmt.Scan(&S)
	fmt.Scan(&S)
	M := make(map[int][]int)
	for i, s := range S {
		ns, _ := strconv.Atoi(string(s))
		M[ns] = append(M[ns], i)
	}
	var ans, l, r int
	for i := 0; i < 10; i++ {
		if m, ok := M[i]; ok {
			l = m[0]
			for j := 0; j < 10; j++ {
				if m, ok := M[j]; ok {
					r = m[len(m)-1]
					for k := 0; k < 10; k++ {
						if m, ok := M[k]; ok && r > l {
							ans += bi(m, l, r)
						}
					}
				}
			}
		}
	}
	fmt.Println(ans)
}
