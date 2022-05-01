/*
    Should use DFS when we are simply trying to find a yes/no answer
    Should use BFS when looking for shortest path
    
    But BFS works in finding a yes/no answer too
    
    approach: BFS
    - We have connected components
    - We are given a start position
    - We are given a destination position
    - The ball can be kicked to any 4 dirs ( up, down, left, right )
    - When the ball is kicked in a direction
        - The ball cannot stop until it hits a wall or hits the edge of the maze
        - So which means our childs to enqueue are not the immediate childs but the childs were its possible for the ball to stop
        - Although we will be used immediate childs to figure out whether we can even go in this direction or not..
    - How do we make sure we do not revisit the same cell thats is already being processed?
        - Once we enqueue a cell, we will mark it visited by changing its value to -1
    
    Time: 
*/

// func hasPath(maze [][]int, start []int, destination []int) bool {
//     m := len(maze)
//     n := len(maze[0])
//     sr := start[0]
//     sc := start[1]
//     dr := destination[0]
//     dc := destination[1]
//     maze[sr][sc] = -1
    
//     dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    
//     q := [][]int{{sr,sc}}
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
        
//         cr := dq[0]
//         cc := dq[1]
//         if cr == dr && cc == dc {return true}
        
//         for _, dir := range dirs {
//             r := cr + dir[0]
//             c := cc + dir[1]
//             for r >= 0 && r < m && c >= 0 && c < n && maze[r][c] != 1 {
//                 r += dir[0]
//                 c += dir[1]
//             }
//             r -= dir[0]
//             c -= dir[1]
//             if maze[r][c] == 0 {
//                 maze[r][c] = -1 // mark this visited
//                 q = append(q, []int{r,c})
//             }
//         }
//     }
    
//     return false
// }



// approach: DFS
func hasPath(maze [][]int, start []int, destination []int) bool {
    m := len(maze)
    n := len(maze[0])
    sr := start[0]
    sc := start[1]
    dr := destination[0]
    dc := destination[1]
  
    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    
    var dfs func(r, c int) bool
    dfs = func(r, c int) bool {
        // base
        if maze[r][c] == -1 {
            return false
        }
        if r == dr && c == dc {return true}
        
        // logic
        // mark this cell visited
        maze[r][c] = -1
                
        for _, dir := range dirs {
            newR := r + dir[0]
            newC := c + dir[1]
            for newR >= 0 && newR < m && newC >= 0 && newC < n && maze[newR][newC] != 1 {
                newR += dir[0]
                newC += dir[1]
            }
            newR -= dir[0]
            newC -= dir[1]
            if reached := dfs(newR,newC); reached {
                return true
            }
            
        }
        return false
    }
    
    return dfs(sr,sc)
}