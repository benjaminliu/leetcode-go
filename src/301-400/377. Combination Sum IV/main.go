package main

func combinationSum4(nums []int, target int) int {
	maps := make(map[int]int)
	count := helper(nums, maps, target)

	return count
}
func helper(nums []int, maps map[int]int, target int) int {

	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}

	if value, ok := maps[target]; ok {
		return value
	}

	count := 0
	for _, value := range nums {
		count += helper(nums, maps, target-value)
	}

	maps[target] = count

	return count
}

//动态规划求解：我们需要一个一维数组dp，其中dp[i]表示目标数为i的解的个数，
// 然后我们从1遍历到target，对于每一个数i，遍历nums数组，如果i>=x, dp[i] += dp[i - x]。
// 这个也很好理解，比如说对于[1,2,3] 4，这个例子，当我们在计算dp[3]的时候，
// 3可以拆分为1+x，而x即为dp[2]，3也可以拆分为2+x，此时x为dp[1]，
// 3同样可以拆为3+x，此时x为dp[0]，我们把所有的情况加起来就是组成3的所有情况了
func combinationSum41(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := range nums {
			if nums[j] <= i {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
}

func main() {

}
