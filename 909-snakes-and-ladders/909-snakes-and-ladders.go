func snakesAndLadders(board [][]int) int {
    n := len(board)
    flatBoard := make([]int, n*n)
    
    right := true
    r := n-1
    c := 0
    i := 0
    for i < n*n {
        flatBoard[i] = board[r][c]
        if flatBoard[i] > 0 {
            flatBoard[i] = flatBoard[i]-1
        }
        i++
        if right {
            c++
            if c == n {
                c--
                r--
                right=false
            }
        } else {
            c--
            if c == -1 {
                c++
                r--
                right=true
            }
        }
    }
    level := 0
    q := []int{0}
    flatBoard[0] = -2
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq == n*n-1 {
                return level
            }
            for i := 1; i <= 6; i++ {
                nIdx := dq + i
                if nIdx < n*n && flatBoard[nIdx] != -2 {
                    if flatBoard[nIdx] == -1 {
                        q = append(q, nIdx)
                    } else {
                        q = append(q, flatBoard[nIdx])
                    }
                    flatBoard[nIdx] = -2
                }
            }
            qSize--
        }
        level++
    }
    return -1
}