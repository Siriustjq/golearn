package main

func main() {

}

func maxUncrossedLines(nums1, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i, v := range nums1 {
		for j, k := range nums2 {
			if v == k {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
