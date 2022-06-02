func findMaxLength(nums []int) int {
    countIdx := map[int]int{0:-1}
    count := 0
    max := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            count++
        } else {
            count--
        }
        idx,seen := countIdx[count]
        if seen && i-(idx+1)+1 > max {
            max = i-(idx+1)+1
        }
        if !seen {
            countIdx[count]=i
        }
    }
    return max
}