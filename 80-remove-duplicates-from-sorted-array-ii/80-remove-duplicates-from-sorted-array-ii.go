func removeDuplicates(nums []int) int {
    slow := 0
    fast := 1
    count := 1
    k := 2
    n := len(nums)
    for fast < n {
        if nums[fast] == nums[fast-1] {
            count++
        } else {
            count = 1
        }
        if count <= k {
            slow++
            nums[slow] = nums[fast]
        }
        fast++
    }
    return slow+1
}