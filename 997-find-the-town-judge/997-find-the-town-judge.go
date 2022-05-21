func findJudge(n int, trust [][]int) int {
    indegrees := make([]int, n+1) // count number of incoming edges to a node
    
    for _, ele := range trust {
        person := ele[0]
        trusts := ele[1]
        
        indegrees[trusts]++
        indegrees[person]--
    }
    
    for i := 1; i < len(indegrees); i++ {
        if indegrees[i] == n-1 {
            return i
        }
    }
    return -1
}