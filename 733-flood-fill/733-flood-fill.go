
/*
    Connected components  == BFS/DFS
    - We need starting point(s) -- all independant nodes 
    There is a domino effect of changing 1 cell color to new color 
    """
        plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, 
        plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
    """
    
    approach : using BFS
    - We are given a starting point ( sr, sc )
    - Enqueue that starting
    - Save the color we are changing from ( oldColor, so we can compare with its neighbors )
    - Change the sr,sc to new color
    - Start processing the queue
    - As soon as we add an item to our queue, we will change the color right away to newColor
    - So that we do not visit this node again
    - if we are not allowed to modify the matrix ( save the visited poins in a set )
    time: o(mn)
    space: o(mn) - for the queue
    

*/




// dfs
// func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    
//     if image == nil {return image}
//     oldCol := image[sr][sc]
//     m := len(image)
//     n := len(image[0])
            
//     var dfs func(r,c,nc,oc int)
//     dfs = func(r,c,nc,oc int) {        
//         // base
//         if r < 0 || r >= m || c < 0 || c >= n || image[r][c] == nc || image[r][c] != oc {
//             return
//         }
        
//         // logic
//         image[r][c] = nc
//         dfs(r+1, c, nc,oc)
//         dfs(r-1, c, nc,oc)
//         dfs(r, c-1, nc,oc)
//         dfs(r, c+1, nc,oc)
//     }
    
//     dfs(sr, sc, newColor, oldCol)
//     return image
// }


func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image == nil {
        return image
    }
    oldCol := image[sr][sc]
    if image[sr][sc] == newColor {return image}  
    m := len(image)
    n := len(image[0])

    q := [][]int{{sr,sc}}
    dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        image[dq[0]][dq[1]] = newColor
        
        for _, dir := range dirs {
            r := dq[0] + dir[0]
            c := dq[1] + dir[1]
            if r >= 0 && r < m && c >= 0 && c < n && image[r][c] == oldCol {
                image[r][c] = newColor
                q = append(q, []int{r,c})
            }
        }
    }
    return image
}
