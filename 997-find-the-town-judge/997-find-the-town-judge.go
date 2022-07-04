func findJudge(n int, trust [][]int) int {
    
    indegrees := make([]int, n+1)
    
    for i := 0; i < len(trust); i++ {
        indegrees[trust[i][1]]++
        indegrees[trust[i][0]]--
    }
    
    for i := 1; i < len(indegrees); i++ {
        if indegrees[i] == n-1 {
            return i
        }
    }
    return -1
    
}