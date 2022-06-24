func hasPath(maze [][]int, start []int, destination []int) bool {
    sr := start[0]
    sc := start[1]
    dr := destination[0]
    dc := destination[1]
    m := len(maze)
    n := len(maze[0])
    
    if maze[dr][dc] == 1 {return false}
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if r < 0 || r == m || c < 0 || c == n || maze[r][c] == 2 {
            return false
        }
        if r == dr && c == dc {return true}
        
        // logic
        maze[r][c] = 2 // mark it visited
        for _, dir := range dirs {
            nr := r + dir[0]
            nc := c + dir[1]
            for nr >= 0 && nr < m && nc >= 0 && nc < n && maze[nr][nc] != 1 {
                nr += dir[0]
                nc += dir[1]
            }
            nr -= dir[0]
            nc -= dir[1]
            reached := dfs(nr, nc)
            if reached {
                return true
            }
        }
        return false
    }
    return dfs(sr, sc)
}