func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    
    max := 0
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        // is this number start of a sequence?
        _, exists := set[num-1];
        // if there is no prev number than this number, than this number is the start of sequence
        // try to build the sequence with this number as long as possible by adding +1 to num and checking in set
        if !exists {
            count := 1
            next := num+1
            _, setContainsNext := set[next]
            for setContainsNext {
                count++
                next++
                _, setContainsNext = set[next]
            }
            if count > max {
                max = count
            }
        }
    }
    return max
}