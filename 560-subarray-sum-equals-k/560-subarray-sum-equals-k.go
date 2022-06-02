func subarraySum(nums []int, k int) int {
    seen := map[int]int{0:1}
    sum := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        complement := sum - k
        count += seen[complement]
        seen[sum]++
    }
    return count
}