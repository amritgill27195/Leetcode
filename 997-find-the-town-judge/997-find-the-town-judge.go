/*
    approach: 
    
    indegrees array to track number of incoming edge at a node,
    then whichever node has exactly n-1 value, thats the node with the most amount
    of incoming egdes 
    
    if we do not find a n-1 value, then that means there is no trusted person, return -1
    
    remember indegrees array counts the number of incoming edges to $this node ( i.e this translates to idx which is a node )
    
    in this case, we will track number of trusts given to a node 
    then look for a node with exactly n-1 value
    
    time: o(n)
    space: o(n) -- for indegrees array
*/

func findJudge(n int, trust [][]int) int {
    indegrees := make([]int, n+1)
    outdegrees := make([]int, n+1)
    
    for _, ele := range trust {
        person := ele[0]
        trusts := ele[1]
        
        indegrees[trusts]++
        outdegrees[person]++
    }
    
    for i := 1; i < len(indegrees); i++ {
        if indegrees[i] == n-1 && outdegrees[i] == 0 {
            return i
        }
    }
    return -1
}
