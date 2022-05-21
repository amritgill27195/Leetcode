func addOperators(num string, target int) []string {
    result := []string{}
    
    var backtrack func(path string, startIdx int, calc int, tail int)
    backtrack = func(path string, startIdx int, calc int, tail int) {
        // base
        if startIdx == len(num) && calc == target {
            result = append(result, path)
            return
        }
        
        // logic
        for i := startIdx; i < len(num); i++ {
                        if string(num[startIdx]) == "0" && startIdx != i {continue}
            currStr := string(num[startIdx:i+1])
            currStrNum, _ := strconv.Atoi(currStr)
            if startIdx == 0 {
                backtrack(path+currStr, i+1, currStrNum, currStrNum)
            } else {
                backtrack(path+"+"+currStr, i+1, calc+currStrNum, currStrNum)
                backtrack(path+"-"+currStr, i+1, calc-currStrNum, -currStrNum)
                backtrack(path+"*"+currStr, i+1, calc-tail+(tail*currStrNum), tail*currStrNum)
            }
        }
    }
    backtrack("",0,0,0)
    return result
}