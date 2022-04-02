
/*
    connected components
    BFS
    We will mark a land visited by changing the land number to 0 so its not recounted by some other child
    
    need to figure out how to do dfs
*/

func maxAreaOfIsland(grid [][]int) int {
    
    
    if grid == nil {
        return 0
    }
    dirs := [][]int{ {1,0}, {-1,0}, {0,1}, {0,-1} }
    m := len(grid)
    n := len(grid[0])
    
    q := [][]int{}
    maxArea := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            
            if grid[i][j] == 1 {
                grid[i][j] = 0
                q = append(q, []int{i,j})
                area := 0
                for len(q) != 0 {
                    area++
                    dq := q[0]
                    q = q[1:]
                    for _, dir := range dirs {
                        r := dq[0] + dir[0]
                        c := dq[1] + dir[1]
                        if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1 {
                            grid[r][c] = 0
                            q = append(q, []int{r,c})
                        }
                    }
                }
                if area > maxArea {
                    maxArea = area
                }
                
            }
        }
    }
    
    return maxArea
}