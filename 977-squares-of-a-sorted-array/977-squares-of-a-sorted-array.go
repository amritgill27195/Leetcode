 /*
    approach 1
    square them all into a new array
    and sort the new arry
    time: o(nlogn)
    space: o(1) if we do not consider output array as part of space
    
    approach 2 : using two pointers
    Have a left and right pointer ( 0, len(nums)-1)
    Cretae an output array of same size
    Have a pointer at len(output)-1
    
    We will sqaure each left and right pointer from nums
    and whoever is bigger will be placed in the back of the output array.
    Whoever was bigger, we will move that pointer
    and we will also move the output ptr back 1
    
    
    we will do this until leftPtr <= rightPtr
    
   
    */


// square each number and store into a new array
// sort the new array and return the new array
// time: o(nlogn)
// space: o(1)
func sortedSquares(nums []int) []int {
    out := []int{}
    for i := 0; i < len(nums); i++ {
        out = append(out, nums[i] * nums[i])
    }
    sort.Ints(out)
    return out
}

// square them in place and sort them at the end
// time: o(nlogn)
// space: o(1)
// func sortedSquares(nums []int) []int {
//     for i := 0; i < len(nums); i++ {
//         nums[i] = nums[i] * nums[i]
//     }
//     sort.Ints(nums)
//     return nums
// }