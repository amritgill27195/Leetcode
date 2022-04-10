func countArrangement(n int) int {
    nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}
    
    count := 0
	var dfs func(start int)
	dfs = func(start int) {
		// base
		if start >= n {
			// fmt.Println(nums)
			count++
		}

		// logic
		for i := start; i < len(nums); i++ {
			// check
			if (start+1)%nums[i] == 0 || nums[i]%(start+1) == 0 {
				// action
				nums[start], nums[i] = nums[i], nums[start]
				// recurse
				dfs(start + 1)
				// backtrack
				nums[start], nums[i] = nums[i], nums[start]
			}
		}
	}
	dfs(0)
	return count
}