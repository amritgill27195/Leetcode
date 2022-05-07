func subarraySum(nums []int, k int) int {
    rsCount := map[int]int{0:1}
    count := 0
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        complement := sum - k
        count += rsCount[complement]
        rsCount[sum]++
    }
    return count
}