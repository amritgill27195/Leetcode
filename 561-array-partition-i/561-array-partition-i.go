/*
    approach: Form and evaluate every single permutation 
    - Using DFS+Backtracking, form every single permutation
    - When we have went over every single num in a path
    - This is when our pivot/start ptr will be out of bounds
    - this will mean that we have a permutation to evalutate
    - we can do our evaluation in the base case
    - and check whether min from a pair sum to a number
    - if this permutation sum > maxSum, save it
    - return maxSum at the end
    
    time: exponential but I cant figure out by how much... is it 2^n!
    space: o(n!)
    Result: TLE
*/
// func arrayPairSum(nums []int) int {
//     maxSum := 0
//     var backtrack func(start int)
//     backtrack = func(start int) {
//         // base
//         if start == len(nums) {
//             sum := 0
//             for i := 0; i < len(nums); i+=2 {
//                 sum += int(math.Min(float64(nums[i]), float64(nums[i+1])))
//             }
//             if sum > maxSum {
//                 maxSum = sum
//             }
//             return
//         }
        
//         // logic
//         for i := start; i < len(nums); i++ {
//             // action
//             nums[start],nums[i] = nums[i],nums[start]
//             // recurse
//             backtrack(start+1)
//             // backtrack
//             nums[start],nums[i] = nums[i],nums[start]
//         }
//     }
//     backtrack(0)
//     return maxSum
// }


/*
    approach: Sort + Linear scan
    - We can sort the array
    - So that the bigger numbers are paired together
    - Thereby increasing our chances of getting a max sum when picking min from bigger pairs
    
    time:
        sort : o(nlogn)
         +
        linear scan: o(n)
        o(nlogn)
    space: o(1)
*/
func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    sum := 0
    for i := 0; i < len(nums); i+=2 {
        sum += int(math.Min(float64(nums[i]), float64(nums[i+1]))) 
    }
    return sum
}
