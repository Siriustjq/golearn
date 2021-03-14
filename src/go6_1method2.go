package main

import "fmt"

func main() {
	tjq := make(map[int]int)
	tjq[0] = 9896
	tmp2 := tjq
	a := 4
	b := a
	a = 7
	fmt.Println(a, " ", b)

	fmt.Println(tmp2[0])
	test1(tjq)
	fmt.Println(tjq[0])
	fmt.Println(test2(tjq)[0])
	fmt.Println(tmp2[0])
}

func test1(m map[int]int) {
	tmp := make(map[int]int)
	tmp = m
	tmp[0] = 857
}

func test2(m map[int]int) map[int]int {
	m[0] = 857
	return m
}
