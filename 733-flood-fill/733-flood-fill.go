
/*
    Connected components  == BFS/DFS
    
    approach : using BFS
    - We need ALL the starting independent nodes
    - Luckily we are given a starting point ( sr, sc )
    - Enqueue that starting
    - Save the color we are changing ( so we can compare with its neighbors )
    - Change the sr,sc to new color
    - Start processing the queue
    - As soon as we add an item to our queue, we will change the color right away to newColor
    - So that we do not visit this node again
    - if we are not allowed to modify the matrix ( save the visited poins in a set )
    time: o(mn)
    space: o(mn) - for the queue
    
    - Problems I ran into with this approach:
        - If the newColor == startingColor, and neighbors are already newColor, you will end up in TLE adding back and forth the same neighbor and parent
        - To mark a neighbor || parent visited -- mark it with a -newColor and change it back at the end to newColor to avoid this.
        

*/


// using level order with BFS
// func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    
//     if image == nil || len(image) == 0 {
//         return image
//     }
//     newColor *= -1
//     queue := [][]int{ {sr,sc} } // [ [r,c] ]
//     oldColor := image[sr][sc]
//     dirs := [][]int{ {1,0},{-1,0},{0,1},{0,-1} }
//     image[sr][sc] = newColor
//     m := len(image)
//     n := len(image[0])
    
//     for len(queue) != 0 {
//         dq := queue[0]
//         queue = queue[1:]
        
//         for _, dir := range dirs {
//             r := dq[0] + dir[0]
//             c := dq[1] + dir[1]
//             if r >= 0 && r < m && c >= 0 && c < n && image[r][c] == oldColor {
//                 image[r][c] = newColor
//                 queue = append(queue, []int{r,c})
//             }
//         }   
//     }
    
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if image[i][j] == newColor {
//                 image[i][j] *= -1
//             }
//         }
//     }
//     return image
    
// }
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image == nil || len(image) == 0 {
        return image
    }

    oldColor := image[sr][sc]
    if oldColor != newColor {
    
    
        var dfs func(img [][]int, r, c, oc, nc int)
        dfs = func(img [][]int, r, c, oc, nc int){

            // base
            // boundary checks
            if img == nil || r < 0 || r >= len(img) || c < 0 || c >= len(img[0])  {return}
            // color checks
            if img[r][c] != oc {return}

            // logic
            img[r][c] = nc
            dfs(img, r+1,c,oc, nc)
            dfs(img, r-1,c,oc, nc)
            dfs(img, r,c+1,oc, nc)
            dfs(img, r,c-1,oc, nc)    
        }
        dfs(image, sr, sc, oldColor, newColor)
    }
        
    return image
}
