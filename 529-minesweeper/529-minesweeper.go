func updateBoard(board [][]byte, click []int) [][]byte {
    
    
    sr := click[0]
    sc := click[1]
    if board[sr][sc] == 'M' {
        board[sr][sc] = 'X'
        return board
    }
    
    board[sr][sc] = 'V' // visited
    q := [][]int{{sr,sc}}
 
    for len(q) != 0 {
        // process current cell
        dq := q[0]
        q = q[1:]
        currRow := dq[0]
        currCol := dq[1]
        
        // check for adjacent mines
        numMines, univistedCells := countMinesAndReturnUnVisitedCells(currRow, currCol, board)
        if numMines > 0 {
            board[currRow][currCol] = byte(numMines + '0')
        } else {
            board[currRow][currCol] = 'B'
            // enqueue univistedCells and mark them visited 
            for _, pos := range univistedCells {
                r := pos[0]
                c := pos[1]
                board[r][c] = 'V'
                q = append(q, pos)
            }
        }   
    }
    
    return board
}
var dirs = [][]int{
        {-1,0}, // up
        {1,0}, // down
        {0,-1}, // left
        {0,1}, //right
        {-1,-1}, // diag-up-left
        {-1,1}, // diag-up-right
        {1,-1}, // diag-down-left
        {1,1}, // diag-down-right
    }

func countMinesAndReturnUnVisitedCells(r,c int, board [][]byte) (int, [][]int) {
    unvisitedCells := [][]int{}
    numMinesFound := 0
    m := len(board)
    n := len(board[0])
    for _, dir := range dirs {
        newR := r + dir[0]
        newC := c + dir[1]

        if newR >= 0 && newR < m && newC >= 0 && newC < n && (board[newR][newC] == 'M' || board[newR][newC] == 'E') {
            if board[newR][newC] == 'M' {
                numMinesFound++
            } else {
                unvisitedCells = append(unvisitedCells, []int{newR,newC})

            }
        }
    }
    return numMinesFound, unvisitedCells
}
