package main

import "fmt"

func main() {
	n := 1
	start := 7
	fmt.Println(xorOperation(n, start))
}

func xorOperation(n int, start int) int {
	var counts []int
	counts = make([]int, n)
	var res int
	for i := 0; i < n; i++ {
		counts[i] = start + 2*i
		if i == 0 {
			res = counts[i]
		} else {
			res = res ^ counts[i]
		}
	}
	return res
}
