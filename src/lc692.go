package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []string{"sr", "saa", "sr", "tjq", "tjq", "ztxm"}
	k := 2
	fmt.Println(topKFrequent(test, k))
}

func topKFrequent(words []string, k int) []string {
	index := make(map[string]int)
	res := make([]string, 0, len(words))
	for _, w := range words {
		index[w]++
	}
	for str := range index {
		res = append(res, str)
	}
	sort.Slice(res, func(i, j int) bool {
		f, l := res[i], res[j]
		return index[f] > index[l] || index[f] == index[l] && f < l
	})
	return res[:k]
}
