/**
    Each value in the array list represents the jump distance we can jump ( 1 thru N where N is the ith val in array list)
    
    approach: DFS, explore all paths
    - We can start from idx 0
    - suppose the value at idx 0 is N
    - Then we explore whether jumping 1..N is possible
    - if any one returns true, we will return true and stop exploring
    
    time: exponential ( n^N - i think, there are n elements each have their own distince N value )
    space: exponential - i think
    Result: TLE
    
**/
// func canJump(nums []int) bool {
//     var dfs func(startIdx int) bool
//     dfs = func(startIdx int) bool {
//         // base
//         if startIdx >= len(nums)-1 {
//             return true
//         }
        
//         // logic
//         for i := nums[startIdx]; i >= 1; i-- {
//             nextIdx := i + startIdx
//             if dfs(nextIdx) {
//                 return true
//             }
//         }
//         return false
//     }
//     return dfs(0)
// }



/**
    Each value in the array list represents the jump distance we can jump ( 1 thru N where N is the ith val in array list)
    
    approach: DFS, explore all paths with Memoization
    - We can try caching in a map[idx]bool
        - why not an array ?
        - well we cannot safetly assume array size because a jump value could be 1k but array size is only 5.. and then it we go out of bounds..
        - therefore map
        
    - We can start from idx 0
    - suppose the value at idx 0 is N, then jump (0+N..1)
    - if this new idx position that we will jump to is already explored ( cached ) and if its true, then return true
    - otherwise recursively call the dfs function again and cache its value - if the value is true, return and exit early, otherwise continue exploring...
    
    time: exponential ( n^N - i think, there are n elements each have their own distince N value )
    space: exponential - i think
    Result: Still TLE
    
**/
func canJump(nums []int) bool {
    var dfs func(startIdx int) bool
    memo := map[int]bool{}

    dfs = func(startIdx int) bool {
        // base
        if startIdx >= len(nums)-1 {
            return true
        }
        
        // logic
        for i := nums[startIdx]; i >= 1; i-- {
            nextIdx := i + startIdx
            if res, ok := memo[nextIdx]; ok {
                if res {return true}
            }else {
                res := dfs(nextIdx)
                memo[nextIdx] = res
                if res { return true }
            }
        }
        return false
    }
    return dfs(0)
}