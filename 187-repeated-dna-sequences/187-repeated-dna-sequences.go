func findRepeatedDnaSequences(s string) []string {
    
    /*
        find all 10 len substrings that are repeated more than once
        
        sliding window ( without the left pointer ... that wont be needed )
        fixed window size -- 10 len
        How do we keep track of window state to be compared later?
        - We will keep a window state and we want to do quickly search
        - therefore a map[string]bool{} -- bool as a flag to dedupe results
        
        The approach is that we will keep looping over the string and forming a temp string.
        as soon as the len of temp string == 10,
        we will check in our window state whether we have seen this or not.
        If we have not, add it. and reset the tmp to "" and move one.
        If we have, then check if the value is false, if it is, save this to output array and update the value of this key to be true (as in successfully saved so no one else re-uses this value again and adds a dupe entry in our output array )
        
    
        time: o(n)
        space: o(n)
    */
    if len(s) <= 1 {
        return nil
    }
    out := []string{}
    // left := 0
    tmp := ""
    windowState := map[string]bool{}
    for i := 0; i < len(s); i++ {
        char := string(s[i])
        tmp += char
        if len(tmp) == 10 {
            // leftChar := string(s[left])
            used, seen := windowState[tmp]
            if !seen {
                windowState[tmp] = false
            } else if seen && !used {
                out = append(out, tmp)
                windowState[tmp] = true
            }
            tmp = string(tmp[1:])
        }
    }
    return out
    
}