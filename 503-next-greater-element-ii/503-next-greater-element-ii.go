func nextGreaterElements(nums []int) []int {
    out := make([]int, len(nums))
    for idx, _ := range out {
        out[idx] = -1
    }
    st := []int{}
    n := len(nums)
    for i := 0 ; i < 2 * n; i++ {
        for len(st) != 0 && nums[i%n] > nums[st[len(st)-1]] {
            top := st[len(st)-1];
            st = st[:len(st)-1]
            out[top] = nums[i%n]
        }
        if i < n {
            st = append(st, i)
        }
    }
    
    return out
}