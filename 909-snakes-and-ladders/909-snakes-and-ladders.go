func snakesAndLadders(board [][]int) int {
    n := len(board)
    flatBoard := make([]int, n*n)
    idx := 0
    r := n-1
    c := 0
    even := 0
    for idx < n*n {
        if board[r][c] == -1 {
            flatBoard[idx]=-1
        } else {
            flatBoard[idx] = board[r][c]-1
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
    
    q := []int{0}
    flatBoard[0] = -2
    count := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == n*n-1 {
                return count
            }
            for i := 1; i <= 6; i++ {
                childIdx := dq+i
                if childIdx < n*n {
                    if flatBoard[childIdx] == -2 {continue}
                    if flatBoard[childIdx] == -1 {
                        q = append(q, childIdx)
                    } else {
                        q = append(q, flatBoard[childIdx])
                    }
                    flatBoard[childIdx] = -2
                }
            }
            qSize--
        }
        count++
    }
    return -1
}