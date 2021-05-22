package main

func main() {

}

func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	begin := 0
	for _, number := range nums {
		begin = begin ^ number
	}
	return begin == 0
}
