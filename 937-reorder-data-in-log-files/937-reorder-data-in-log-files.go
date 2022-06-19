func reorderLogFiles(logs []string) []string {
    sort.SliceStable(logs, func(i, j int) bool{
        
        iLog := logs[i]
        jLog := logs[j]
        iLogSplit := strings.SplitN(iLog," ", 2) // "dig1 8 1 5 1" = ["dig1", "8 1 5 1"]
        jLogSplit := strings.SplitN(jLog," ", 2) // "g1 act car" = ["g1", "act car"]
        iLogIsDigitLog := isDigit(iLogSplit[1][0])
        jLogIsDigitLog := isDigit(jLogSplit[1][0])
        
        if !iLogIsDigitLog && !jLogIsDigitLog { // both are letter logs
            
            /*
                ["g2", "act car"]
                ["g1", "act car"]
            */
            if iLogSplit[1] == jLogSplit[1] {
                return iLogSplit[0] < jLogSplit[0]
            } else {
                // res := strings.Compare(iLogSplit[1],jLogSplit[1])
                // if
                return iLogSplit[1] < jLogSplit[1]
            }
            
            
        } else if iLogIsDigitLog && !jLogIsDigitLog { // i is digit and j is letter, letter comes first, so swap
            return false
        } else if !iLogIsDigitLog && jLogIsDigitLog { // i is letter and j is digit, letter comes first, so do nothing
            return true
        } else { // both are digit log, nothing to swap, return false
            return false
        }        
    }) 

    return logs
}

func isDigit(n byte) bool {
    return n >= '0' && n <= '9'
}