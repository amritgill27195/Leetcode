func addOperators(num string, target int) []string {
    
    var result []string
    
    var dfs func(path string, start, calc, tail int)
    dfs = func(path string, start, calc, tail int) {
        // base
        if start == len(num) && calc == target {
            result = append(result, path)
            return
        }
        
        for i := start; i < len(num); i++ {
            if string(num[start]) == "0" && start != i {continue}
            currStr := string(num[start:i+1])
            currStrNum, _ := strconv.Atoi(currStr)
            if start == 0 {
                dfs(path+currStr, i+1, currStrNum, currStrNum)
            } else {
                dfs(path+"+"+currStr, i+1, calc+currStrNum,currStrNum)
                dfs(path+"-"+currStr, i+1, calc-currStrNum,-currStrNum)
                dfs(path+"*"+currStr, i+1, calc-tail+(tail*currStrNum), tail*currStrNum)
            }
        }
    }
    dfs("", 0,0,0)
    return result
}