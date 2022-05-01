
/*
    We are not looking for shortest path, simply count areas of an island
    

    approach: DFS
    - For each 1 we run into
    - deploy dfs from that r,c
        - DFS will mark each node visited by changing a land(1) to water(0)
        - This way the same cell is not re-added by a neighboring land
    - The dfs will return the area of the island formed using the r,c from above
    - Compare and save the returned area with max area
*/

func maxAreaOfIsland(grid [][]int) int {
    max := 0
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] == 0 {
            return 0
        }
        
        // logic
        count := 1
        grid[r][c] = 0
        for _, dir := range dirs {
            count += dfs(r+dir[0], c+dir[1])
        }
        return count
    }
    
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                area := dfs(i,j)
                if area > max {
                    max = area
                }
            }
        }
    }
    
    return max
}