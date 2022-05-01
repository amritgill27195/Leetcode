
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


// DFS ( recursive )
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    oldColor := image[sr][sc]
    if newColor == oldColor {
        return image
    }
    m := len(image)
    n := len(image[0])
    
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    var dfs func(r, c int) 
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || image[r][c] == newColor || image[r][c] != oldColor {
            return
        }
        
        // logic
        image[r][c] = newColor
        for _, dir := range dirs {
            dfs(r+dir[0],c+dir[1])
        }
    }
    dfs(sr, sc)
    return image
}

