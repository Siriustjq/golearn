package main

import (
	"fmt"
)

func main() {
	numsTjq := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	for _, num := range findOrder(4, numsTjq) {
		fmt.Println(num)
	}
}

func findOrder(numCrouse int, prerequisites [][]int) []int {
	const (
		NOTVISIT = iota
		VISITING
		VISITED
	)
	hasRing := false
	dfsMap := make(map[int][]int)
	hasMap := make(map[int]bool)
	for i := 0; i < len(prerequisites); i++ {
		dfsMap[prerequisites[i][0]] = append(dfsMap[prerequisites[i][0]], prerequisites[i][1])
		hasMap[prerequisites[i][0]] = true
		hasMap[prerequisites[i][1]] = true
		//fmt.Printf("key is %d, value is %d",prerequisites[i][0], prerequisites[i][1])
	}
	visited := make(map[int]int) //标记某元素是否使用过 在有向图无环的时候需要配置三个量：未访问、访问中，已访问
	var searchDfs func(nums []int)
	var order []int
	searchDfs = func(nums []int) {
		for _, num := range nums {
			if hasRing {
				return
			} else {
				if visited[num] == NOTVISIT {
					visited[num] = VISITING
					searchDfs(dfsMap[num])
					order = append(order, num)
					visited[num] = VISITED
				} else if visited[num] == VISITING {
					hasRing = true
					return
				}
			}
		}
	}
	var keys []int
	//注意返回值
	for key := range dfsMap {
		//fmt.Println(key)
		keys = append(keys, key)
	}
	searchDfs(keys)
	if hasRing {
		res := []int{}
		return res
	}
	if len(order) < numCrouse {
		res := []int{}
		for i := 0; i < numCrouse; i++ {
			if !hasMap[i] {
				res = append(res, i)
			}
		}
		res = append(res, order...)
		return res
	}
	return order
}
