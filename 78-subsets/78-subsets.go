/*
    approach 1: 0/1 recursion
    - choose / not choose
    - coding out the choose / not choose decision tree
    - we can only use the 1 num ONCE
    - At leaves, we will append the running path
    - ( new path each recursion or backtrack existing )
    - time: exponential ( o(2^n) where n is the number of elements in nums array )
    - space: o(h) the max height of our decision tree which looks like it will be len(arr)

    approach 2: for loop based recursion
    - instead of looking for an answer at a leaf node or when target == 0, we dont have such constraints here
    - each path at every node is an answer
    - time: exponential o(2^n)
    - space: o(h) the max height of our decision tree which looks like it will be len(arr)

    approach 3: no recursion
    - Start with [[]] list
    - For each number in nums,
        - For each list in result,
            create a new list with curr num + eachList contents
    - time: exponential o(n*2^n)
*/


// 0/1 based recursion
func subsets(nums []int) [][]int {
    var result [][]int
    var dfs func(start int, paths []int)
    dfs = func(start int, paths []int) {
        // base
        if start >= len(nums) {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        
        // logic
        // not choose
        dfs(start+1, paths)
        
        // choose
        paths = append(paths, nums[start])
        dfs(start+1, paths)
        paths = paths[:len(paths)-1]
    }
    dfs(0, nil)
    return result
}