func confusingNumberII(n int) int {
    
    count := 0
    flipMap := map[int]int{
        0:0,1:1,6:9,8:8,9:6,
    }
    var dfs func(curr int)
    dfs = func(curr int) {
        // base
        if curr > n {return}
        
        // logic
        if isConfusingNum(curr, flipMap) {
            count++
        }
        for k, _ := range flipMap {
            newNum := curr * 10 + k
            if newNum != 0 {
                dfs(newNum)
            }
        }
    }
    dfs(0)
    return count
}

func isConfusingNum(n int, flipMap map[int]int) bool {
    orig := n
    result := 0
    for n != 0 {
        last := n % 10
        n = n / 10
        lastFlipped := flipMap[last]
        result = result * 10 + lastFlipped
    }
    return orig != result
} 