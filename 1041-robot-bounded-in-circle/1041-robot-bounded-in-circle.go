func isRobotBounded(instructions string) bool {
    //                U     R      D     L
    dirs := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    x := 0
    y := 0
    idx := 0
    for i := 0 ; i < len(instructions); i++ {
        inst := instructions[i]
        if inst == 'R' {
            idx = (idx + 1)% 4
        } else if inst == 'L' {
            idx = (idx + 3) % 4
        } else if inst == 'G'{
            x += dirs[idx][0]
            y += dirs[idx][1]
        }
        
    }
    
    return idx != 0 || (x == 0 && y == 0)
}