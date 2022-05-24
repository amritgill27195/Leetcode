func findMaxLength(nums []int) int {
    count := 0
    max := 0
    countIdx := map[int]int{0:-1}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0{count--}
        if nums[i] == 1{count++}
        idx, seen := countIdx[count]
        idx++
        if !seen {
            countIdx[count] = i
        } else {
            cl := i-idx+1
            if cl > max {
                max = cl
            }
        }
    }
    return max
}