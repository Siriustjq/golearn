package main

import "fmt"

var tool1 map[int]int

func main() {

	tool1 = make(map[int]int)
	tool1[2] = 3
	tool1 = make(map[int]int)
	tool1[2] = 4
	fmt.Println(tool1[2])
}
func justForFun(a int) {
	tool1 = make(map[int]int)
}
