/*
    approach: for loop based recursion
    - We need to output all possible expressions using ALL nums in num string that result into target
    - The only operators we have to try are '+','-','*'
    - Which also means we need make 3 recursive calls for EACH ith substr we have formed.
    - How do we form all the expressions?
        - For loop based recursion.
        - We will have a start idx that will be passed into the initial recursive call
        - And our for loop starts from that start idx until it reaches the end of num string
        - If need to form '1', '12', '123' from the given '123' string,
            - we can create substr using str[startIdx:i] -- i keeps moving and start stays where it is - so that works well in our favor
        - But we need to form expressions...
        - For each ith element, we need to form 3 expressions with '+','-','*' (i.e 3 recursive calls)
    - What are all the parameters we need to pass down in the recursive call ?
    - How do we maintain and enforce PEMDAS operator precedence?
    - What evaluates whether this expression works or not?
    - What do we need to backtrack? 

*/
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