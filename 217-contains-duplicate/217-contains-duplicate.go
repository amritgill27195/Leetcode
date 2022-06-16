func containsDuplicate(nums []int) bool {
    seen := map[int]bool{}
    for i := 0; i < len(nums); i++ {
        _, ok := seen[nums[i]]
        if ok {
            return true
        }
        seen[nums[i]] = false
    }
    return false
}