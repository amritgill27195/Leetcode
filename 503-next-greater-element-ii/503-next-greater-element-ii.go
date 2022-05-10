/*
    approach: brute force
    - nested for loop on doubled array len and use % n to get relative idx within array size
    - the nested for loop goes until 
    time: o(n^2)
    
*/

func nextGreaterElements(nums []int) []int {
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        out[i] = -1
    }
    
    n := len(nums)
    
    for i := 0; i < len(nums); i++ {
        for j := i+1;  j % n != i; j++ {
            if nums[j%n] > nums[i] {
                out[i] = nums[j%n]
                break
            }
        }
    }
    return out
}