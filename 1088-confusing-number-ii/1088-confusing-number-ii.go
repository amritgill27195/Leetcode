func confusingNumberII(n int) int {
    count := 0
    flipMap := map[int]int{0:0,1:1,6:9,8:8,9:6}
    var dfs func(curr int)
    dfs = func(curr int) {
        // base
        if curr > n {return}
        if isConfusing(flipMap, curr) { count++ }
        // logic
        for k, _ := range flipMap {
            newCurr := curr * 10 + k
            if newCurr != 0 {
                dfs(newCurr)
            }
        }
    }
    dfs(0)
    return count
}

func isConfusing(flipMap map[int]int, n int) bool {
    orig := n
    flip := 0
    for n != 0 {
        last := n%10
        n /= 10
        lastFlipped := flipMap[last]
        flip = flip * 10 + lastFlipped
    }
    return flip != orig
}