func addOperators(num string, target int) []string {
    var result []string
    var dfs func(start int,expr string, calc int, tail int)
    dfs = func(start int,expr string, calc int, tail int) {
        // base
        if start == len(num) {
            if calc == target {
                result = append(result, expr)
            }
            return
        }
        
        
        // logic
        for i := start; i < len(num); i++ {
            if string(num[start]) == "0" && start != i {continue}
            
            // we have 3 choices ; + , - , x
            currNumStr := string(num[start:i+1])
            currNumInt, _ := strconv.Atoi(currNumStr)

            if start == 0 {
                dfs(i+1,currNumStr, currNumInt, currNumInt)
            } else {
                // +
                dfs(i+1,expr+"+"+currNumStr,calc+currNumInt, currNumInt)

                // -
                dfs(i+1,expr+"-"+currNumStr,calc-currNumInt, -currNumInt)

                // x
                dfs(i+1,expr+"*"+currNumStr,(calc-tail)+(tail*currNumInt), tail*currNumInt)
            }
        }
        
    }
    dfs(0,"",0,0)
    return result
}