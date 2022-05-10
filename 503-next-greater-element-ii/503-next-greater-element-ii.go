/*
    approach: brute force
    
*/

func nextGreaterElements(nums []int) []int {
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        out[i] = -1
    }
    
    n := len(nums)
    
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < 2 * len(nums) && j % n != i; j++ {
            if nums[j%n] > nums[i] {
                out[i] = nums[j%n]
                break
            }
        }
    }
    return out
}