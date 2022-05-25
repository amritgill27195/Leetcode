/*
    approach: BFS
    - How did we end up with BFS?
        - We have a start position
        - We have a end poistion
        - We have to reach from start to end in minimum number of moves as possible
        - We have 6 childs to explore for each idx ( 1 to 6 )
        - Therefore BFS
    - The tricky part of this BFS is to figure out what direction to go into if a childIdx goes out of bounds
        - We could try and simulate this but will be complicated
        - Instead what we will do is flatten the board into a 1D array ( since n * n)
        - Then fill out the 1D array with the values from board
            - Fillout the way game wants us to walk the board
            - Start from last row , first col and work ourself up and changing direction each row we go up.
        - If at a specific location of our board, we do have a non-negative integer, then in the flatten array for this position
        - We will store that value-1 -- because the board is n*n with the values from 1 to n*n(inclusive)
        - However there are only n*n-1 indicies, so if we are taking a ladder or getting bit by a snake, we need to know in our flatten array, which idx to go to.
        - Luckily if the board at a specific position is asking us to a different cell, that cell value has a idx of cellVal-1
    - Our BFS process will be working off of indicies from the flattened array
        - Initially enqueue the 0th idx
        - We will be taking the size since we are counting minimum number of moves as poissible
        - Minimum number of moves == a level
        - So once a level is done, level++
        - When processing a level, we will a dq'd idx value
        - If this dq'd idx == n*n-1 (i.e destination idx ), simply return level value
        - Otherwise explore this dq'd idx's 6 childs
            - The 6 childs are generated by rolling a dice
            - We could roll a 1,2,3,4,5,6
            - Therefore from the current idx ( dq'd idx ), we will add all possible 6 jumps to it
            - so child = dq'd idx + i(1..6)
            - if a child is already visisted or going out of bounds( >= n*n ) then skip
            - check if at that new cell/childIdx, are we taking a ladder or getting bit by a snake
            - Our flattened array will this answer at flattenedArray[childIdx]
            - If the value in flattened array is -1, then nothing special, just enqueue this childIdx
            - If the value in flattened array is not -1, and is a positive number, then enqueue the value at childIdx
    - If our BFS ends without ever returning, that means its impossible to reach destination, then return -1
    
    Time: o(n*n) -- i.e o(n^2)
    - we do see all cells in the board
    - n is the num of rows and cols
    - we traverse the entire n*n board/matrix
    Space:  o(n*n) -- i.e o(n^2)
    - we flattened the entire board of n*n into n*n 1D array 
*/
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
                if childIdx >= n*n || flatBoard[childIdx] == -2 { continue }
                if flatBoard[childIdx] == -1 {
                    q = append(q, childIdx)
                } else {
                    q = append(q, flatBoard[childIdx])
                }
                flatBoard[childIdx] = -2
            }
            qSize--
        }
        count++
    }
    return -1
}