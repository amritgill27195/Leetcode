func removeDuplicates(nums []int) int {
    count := 1
    slow := 1
    fast := 1
    n := len(nums)
    k := 2
    
    for fast < n {
        if nums[fast] == nums[fast-1] {
            count++
        } else {
            count = 1
        }
        if count <= k {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return slow
}