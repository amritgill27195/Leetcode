// time: o(n)
// space: o(1)
func isRobotBounded(instructions string) bool {
    //                U     R      D     L
    dirs := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    x := 0
    y := 0
    idx := 0
    for i := 0 ; i < len(instructions); i++ {
        inst := instructions[i]
        if inst == 'R' {
            idx = (idx + 1) % 4
        } else if inst == 'L' {
            idx = (idx + 3) % 4
        } else if inst == 'G'{
            x += dirs[idx][0]
            y += dirs[idx][1]
        }
        
    }
    
    return idx != 0 || (x == 0 && y == 0)
}


/*
    initial naive code after understanding the explanation
    time: o(n)
    space: o(1)
*/

// func isRobotBounded(instructions string) bool {
//     x := 0
//     y := 0
//     dir := "u"
//     idx := 0
    
//     for idx < len(instructions) {
//         inst := string(instructions[idx])
//         if inst == "G" {
//             if dir == "u" {
//                 y++
//             } else if dir == "d" {
//                 y--
//             } else if dir == "l" {
//                 x--
//             } else if dir == "r" {
//                 x++
//             }
//         } else if inst == "L" {
//             if dir == "u" {
//                 dir = "l"
//             } else if dir == "l" {
//                 dir = "d"
//             } else if dir == "d" {
//                 dir = "r"
//             } else if dir == "r" {
//                 dir = "u"
//             }
//         } else if inst == "R" {
//             if dir == "u" {
//                 dir = "r"
//             } else if dir == "l" {
//                 dir = "u"
//             } else if dir == "d" {
//                 dir = "l"
//             } else if dir == "r" {
//                 dir = "d"
//             }
//         }
//         idx++
//     }
//     return dir != "u" || (x == 0 && y == 0)
// }