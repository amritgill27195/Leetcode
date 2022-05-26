func isRobotBounded(instructions string) bool {
    x := 0
    y := 0
    dir := "u"
    idx := 0
    
    for idx < len(instructions) {
        inst := string(instructions[idx%len(instructions)])
        if inst == "G" {
            if dir == "u" {
                y++
            } else if dir == "d" {
                y--
            } else if dir == "l" {
                x--
            } else if dir == "r" {
                x++
            }
        } else if inst == "L" {
            if dir == "u" {
                dir = "l"
            } else if dir == "l" {
                dir = "d"
            } else if dir == "d" {
                dir = "r"
            } else if dir == "r" {
                dir = "u"
            }
        } else if inst == "R" {
            if dir == "u" {
                dir = "r"
            } else if dir == "l" {
                dir = "u"
            } else if dir == "d" {
                dir = "l"
            } else if dir == "r" {
                dir = "d"
            }
        }
        idx++
    }
    return dir != "u" || (x == 0 && y == 0)
}