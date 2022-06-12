func gameOfLife(board [][]int)  {
    m := len(board)
    n := len(board[0])
    // 2 = new alive = 1
    // 3 = new dead = 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            liveCount := countLiveNeighbors(i,j,board)
            if board[i][j] == 1 {
                if liveCount < 2 || liveCount > 3 {
                    board[i][j] = 3
                }
            } else if board[i][j] == 0 {
                if liveCount == 3 {
                    board[i][j] = 2
                }
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 2 {board[i][j] = 1}
            if board[i][j] == 3 {board[i][j] = 0}
        }
    }
    
}

func countLiveNeighbors(r, c int, board [][]int) int {
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,-1},
        {0,1},
        {-1,-1},
        {-1, 1},
        {1,-1},
        {1,1},
    }
    m := len(board)
    n := len(board[0])
    count := 0
    for _, dir := range dirs {
        nr := r+dir[0]
        nc := c+dir[1]
        if nr >= 0 && nr < m && nc >= 0 && nc < n && (board[nr][nc] == 1 || board[nr][nc] == 3) {
            count++
        }
    }
    return count
}