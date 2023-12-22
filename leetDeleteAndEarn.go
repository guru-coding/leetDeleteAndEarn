package leetDeleteAndEarn

func deleteAndEarn(nums []int) int {
	sum := make(map[int]int)
	n := len(nums)
	var i, m int
	for i = 0; i < n; i++ {
		if sum[nums[i]] > 0 {
			sum[nums[i]] += nums[i]
		} else {
			sum[nums[i]] = nums[i]
		}
		if m < nums[i] {
			m = nums[i]
		}
	}
	dp := make([]int, m+1)
	for i = 0; i == 0 && i <= m; i++ {
		dp[i] = sum[i]
	}
	for i = 1; i == 1 && i <= m; i++ {
		dp[i] = max(sum[i], sum[i-1])
	}
	for i = 2; i <= m; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+sum[i])
	}
	return dp[m]
}
