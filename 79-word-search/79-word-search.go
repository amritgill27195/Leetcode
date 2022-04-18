/*
    Connected components : BFS or DFS
    Why BFS wont work here:
    - Because we need to backtrack
    - Why do we need to backtrack ?
        - We may have 1 path that lead to partial word but not full
        - If we went down the partial word path, and marked all the nodes visited, then the path that forms the full word, will not go to visited nodes that were seen by the partial word path. Therefore the partial word path that was unsuccessful, must backtrack its visiting action on those nodes so that if there is a potential full word path, we can return the right answer
    - Why cant we backtrack in BFS?
        - Because we do not know which parent enqueued a node.

    approach: DFS with Backtracking
    - We will start by looping grid looking for 0th idx char from word
    - As soon as we find a char on board that == word[0], we deploy our DFS from that row, col
        - as in go explore if this is a potential path
        - which means our DFS needs to return whether this was an valid starting point back to our for loop running on grid..
        - If the dfs returns back true, exit early, no need to go look for another starting point of our DFS
        - or else continue looping the grid and look for another r,c that == word[0] ( potential starting path of our word search )
    - The DFS itself will recursively call itself in all 4 directions with but the DFS needs to know the next char its looking for
    - So the DFS will need to take a idx pointer that points to char we are searching for, as well as newRow, newCol
    - What about backtracking?
        - At each cell, if the value == word[idx], save the current value in recursion stack in a temp var and mark it visited ( change the value to '#' )
        - when the recursion comes back, change the value '#' back to the value saved in the tmp var
    - If our idx for our word, goes out of bounds, this means we have found all the words, return true and exit early
*/

func exist(board [][]byte, word string) bool {
    
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    var backtrack func(r,c,wordIdx int) bool
    backtrack = func(r,c,wordIdx int) bool {
        // base
        if wordIdx == len(word) {
            return true
        }
        if r < 0 || r == m || c < 0 || c == n || board[r][c] == '#' || board[r][c] != word[wordIdx] {
            return false
        }
        
        // logic
        // action ( save value and mark this position visited )
        tmp := board[r][c]
        board[r][c] = '#'
        // recurse = check all 4 dirs for the next word ( wordIdx+1)
        for _, dir := range dirs {
            if backtrack(r+dir[0], c+dir[1], wordIdx+1){
                return true
            }
        }
        // backtrack - undo the visiting action and restore state back in board
        board[r][c] = tmp
        return false
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                found := backtrack(i,j, 0)
                if found { // we found the full word starting from i,j position
                    return true
                }
            }
        }
    }
    return false
}