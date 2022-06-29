/*
    approach: brute force
    - For each 1 we run into, we will try to form the largest square
    - How do we do that?
        - We will do diagonally down right first
        - If the this value is a 1, then we need to check 2 things
            1. From this cell, go up as far as we can in the same col UNTIL we hit a ceiling row - the ceiling in this case is the ith idx
            2. From this cell, go left as far as we can in tge same row UNTIL we hit a ceiling col - the ceiling col is the jth idx
        - If we were able to successfully do this, then continue diagonal down
        - Otherwise explore the next 1 
*/

func maximalSquare(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    max := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                l := 1
                flag := true
                for i + l < m && j + l < n && flag {
                    
                    for k := i+l; k >= i; k--{
                        if matrix[k][j+l] == '0' {
                            flag = false
                            break
                        }
                    }
                    for k := j+l; k >= j; k--{
                        if matrix[i+l][k] == '0' {
                            flag = false
                            break
                        }
                    }
                    
                    if flag { l++ }
                }
                if l*l > max { max = l*l }
            }
        }
    }
    return max
}