

/*
    connected components!
    A dist in a cell (that already reached 0) can tell us a dist for a cell that cannot reach 0 directly
    
    approach: BFS
    - We need starting INDEPENDANT nodes.
    - 0s do not depend on anything, we are not looking for distance of each 0 to another 0
    - So 0s are independant
    - Enqueue all 0 positions
    - While we are all looking for all 0's , we will mark all the 1's = -1
        - Why? Well this is to distinguish between 1 as a distance vs visited/not-visited.
    - So when we are processing all 0's from our queue
    - We will look for all -1's ( unvisited nodes around this 0 )
    - When we find it, we will enqueue all neighbouring -1's ( but change the value of them to currentNode val + 1)
    - Then we will have cells with 1's in our queue
    - Keep looking for all -1's around each item from our queue, keep adding current val + 1 to new all -1's we find.
    - Instead of keeping the distance maintained at each level, the cell being processed itself will tell us how far away this cell is from a 0.
    - We can take that info and add + 1 for its neighbors that are -1.
    - And then repeat until full matrix is processed
    
    Time: o(mn)
    space: o(mn)
*/

func updateMatrix(mat [][]int) [][]int {
    
    q := [][]int{}
    m := len(mat)
    n := len(mat[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                q = append(q, []int{i,j})
            } else {
                mat[i][j] = -1
            }
        }
    }
    
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        for _, dir := range dirs {
            r := dq[0] + dir[0]
            c := dq[1] + dir[1]
            if r >= 0 && r < m && c >=0 && c < n && mat[r][c] == -1 {
                mat[r][c] = mat[dq[0]][dq[1]] + 1
                q = append(q, []int{r,c})
            }
        }
    } 
    return mat
}