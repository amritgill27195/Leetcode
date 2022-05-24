func subarraySum(nums []int, k int) int {
    m := map[int]int{0:1}
    rs := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        rs += nums[i]
        complement := rs-k
        numTimes, ok := m[complement]
        if ok {
            count += numTimes
        }
        m[rs]++
    }
    return count
}