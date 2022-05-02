
// BFS with extra space ( using another matrix to keep track of visited {r:c} )
func updateMatrix(mat [][]int) [][]int {
    q := [][]int{}
    m := len(mat)
    n := len(mat[0])
    visited := make([][]bool, m)
    for idx, _ := range visited {
        visited[idx] = make([]bool, n)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {
                q = append(q, []int{i,j})
                visited[i][j] = true
            }
        }
    }
    dirs := [][]int{{1,0},{-1,0}, {0,-1},{0,1}}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        cv := mat[cr][cc]
        for _, dir := range dirs {
            r := cr+dir[0]
            c := cc+dir[1]
            if r >= 0 && r < m && c >= 0 && c < n && mat[r][c] == 1 && !visited[r][c] {
                visited[r][c] = true
                mat[r][c] = cv + 1
                q = append(q, []int{r,c})
            }
        }
    }
    return mat
}