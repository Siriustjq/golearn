package main

import (
	"fmt"
	"math"
)

type point struct {
	X, Y float64
}

func (p point) Distance(q point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type path []point //任意类型均可以有方法，只要其不是指针类型或者接口类型

func (p path) Distance() float64 {
	sum := 0.0
	for i, _ := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	perim := path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}
