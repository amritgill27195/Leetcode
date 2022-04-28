/*
    Current confusion: When to/not-to use backtracking in BFS/DFS and why we did not use DP tables in previous backtracking/recursion problems where there may have been repeated subproblems..
        - If you see a need for backtracking, then forget BFS, backtracking goes hand in hand with DFS ( for example populating a path to all leaf nodes in a tree )
    
    
    Approach: Top down DP
    - Top Down DP is starting from the bigger problem and going down towards a smaller problem
    - Top Down DP == memoization
    - Usually used with recursion as a cache to signify that "hey we have already solved this recursive case"
    - The robot can only move either down or right at any point in time
    - So our decision tree has two choices
    - Go explore row down or col right

*/

// Top down DP with memoization
func uniquePaths(m int, n int) int {
    
    memo := make([][]int, m)
    for i := 0; i < m; i++ {
        memo[i] = make([]int, n)
    }

    var dfs func(i, j int) int
    dfs = func(i, j int) int{
        // base
        if i == m-1 && j == n-1 {
            return 1
        }
        if i == m || j == n {return 0}
        
        
        // logic
        if memo[i][j] == 0 {
            // we have not solved this repeated subproblem
            memo[i][j] = dfs(i+1,j) + dfs(i,j+1)
        }
        // we have solved this i,j subproblem, returned the cached/memoized result
        return memo[i][j]
    }
    return dfs(0,0)
}



// brute force: DFS/recursion (TLE)
// func uniquePaths(m int, n int) int {
//     var dfs func(i, j int) int
//     dfs = func(i, j int) int{
//         // base
//         if i == m-1 && j == n-1 {
//             return 1
//         }
//         if i == m || j == n {return 0}
        
//         // logic
//         return dfs(i+1,j) + dfs(i,j+1)
//     }
//     return dfs(0,0)
// }




