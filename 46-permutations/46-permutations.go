func permute(nums []int) [][]int {
    
    result := [][]int{}
    var dfs func(start int)
	dfs = func(start int) {
		if start == len(nums) {
            newL := make([]int, len(nums))
            copy(newL, nums)
            result = append(result, newL)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	dfs(0)
	return result    
}