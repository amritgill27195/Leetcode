/*
    
    approach : BFS
    - Connected components, 1 rooms distance from a gate tells its neighbours rooms distance from a gate
    - What are the independent nodes we can start with?
        - Rooms distance depends on the nearest gate avail.
        - Gates do not depend on anything.
        - Therefore gates are independent and can be used as our starting positions
    - From every gate we look for empty rooms in its immediate vicinity
    - If we found a room as a neighbour of our gate being processed, than that means, this room is 1 distance away
        - current node val + 1 ( gate values are 0 )
    - Then this room will be enqueued "as a gate" for its neighbors. If this room is 1 distance away from nearest gate, its neighbours will be 
        currentRoomDist + 1.
    - Then we can use this logic to mark all rooms distance ( like a domino effect )
    - Do we need to mark anything visited?
        - No? because we will only ever enqueue empty rooms into our queue - so if a room is processed, it will no longer be empty.
    

*/


func wallsAndGates(rooms [][]int)  {
    
    if rooms == nil {
        return
    }
    
    m := len(rooms)
    n := len(rooms[0])
    q := [][]int{}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if rooms[i][j] == 0 {
                q = append(q,[]int{i,j}) // enqueue the gate
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
            if r >= 0 && r < m && c >= 0 && c < n && rooms[r][c] == 2147483647 {
                rooms[r][c] = rooms[dq[0]][dq[1]] + 1
                q = append(q, []int{r,c})
            }
        }
        
    }
    
    
}