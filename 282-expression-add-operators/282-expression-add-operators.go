func addOperators(num string, target int) []string {
    
    var result []string
    
    var dfs func(start int, calc, tail int, expr string)
    dfs = func(start int, calc, tail int, expr string){
        // base
        if start == len(num) && calc == target {
            result = append(result, expr)
            return
        }
        
        // logic
        for i := start; i < len(num); i++ {
            if string(num[start]) == "0" && start != i {continue}
            currStr := string(num[start:i+1])
            currInt, _ := strconv.Atoi(currStr)
            if start == 0 {
                dfs(i+1, currInt, currInt, currStr)
            } else {
                dfs(i+1, calc+currInt, currInt, expr+"+"+currStr)
                dfs(i+1, calc-currInt, -currInt, expr+"-"+currStr)
                dfs(i+1, calc-tail+(tail*currInt), tail*currInt, expr+"*"+currStr)
            }
        }
    }
    dfs(0, 0, 0,"")
    return result
}