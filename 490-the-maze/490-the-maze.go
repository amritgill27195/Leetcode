func hasPath(maze [][]int, start []int, destination []int) bool {
    
    m := len(maze)
    n := len(maze[0])
    
    sr := start[0]
    sc := start[1]
    dr := destination[0]
    dc := destination[1]
    
    maze[sr][sc] = -1
    q := [][]int{{sr,sc}}
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        
        cr := dq[0]
        cc := dq[1]
        
        if cr == dr && cc == dc {return true}
        
        for _, dir := range dirs {
            r := cr + dir[0]
            c := cc + dir[1]
            for r >= 0 && r < m && c >= 0 && c < n && maze[r][c] != 1 {
                r += dir[0]
                c += dir[1]
            }
            r -= dir[0]
            c -= dir[1]
            if maze[r][c] == 0 {
                maze[r][c] = -1                
                q = append(q, []int{r,c})   
            }
        }
    }
    
    return false
}