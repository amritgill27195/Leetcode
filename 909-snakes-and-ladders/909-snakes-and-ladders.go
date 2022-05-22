func snakesAndLadders(board [][]int) int {
    n := len(board)
    moves := make([]int, n*n)
    r := n-1
    c := 0
    idx := 0
    even := 0
    for idx < n*n {
        if board[r][c] != -1 {
            moves[idx] = board[r][c]-1
        } else {
            moves[idx] = board[r][c]
        }
        idx++
        if even % 2 == 0 {
            c++
            if c == n {
                c--
                r--
                even++
            }
        } else {
            c--
            if c == -1 {
                c++
                r--
                even++
            }
        }
    }
    
    level := 0
    q := []int{0}
    moves[0] = -2
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == n*n-1 {
                return level
            }
            for i := 1; i <= 6; i++ {
                child := dq + i
                if child < n*n {
                    if moves[child] != -2 {
                        if moves[child] == -1 {
                            q = append(q, child)
                        } else {
                            q = append(q, moves[child])
                        }
                        moves[child] = -2
                    }
                }
            }
            qSize--
        }
        level++
    } 
    return -1
}