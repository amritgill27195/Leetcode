func calculate(s string) int {
    stack := []int{}
    lastOp := "+"
    curr := 0
    for i, char := range s {
        stringChar := string(char)
        n, _ := strconv.Atoi(stringChar)

        if isDigit(stringChar) {
            curr = curr * 10 + n
        }
        if (!isDigit(stringChar) && stringChar != " ")  || i == len(s)-1 {
            // we either have a new Op or we have reached the last number in string
            // process current op before changing to new op
            if lastOp == "+" {
                stack = append(stack, curr)
            } else if lastOp == "-" {
                stack = append(stack, -curr)
            } else if lastOp == "*" {
                // pop top, multiply with curr and push the new result
                stack[len(stack)-1] = stack[len(stack)-1] * curr
            } else if lastOp == "/" {
                // pop top, divide with curr and push the new result
                stack[len(stack)-1] = stack[len(stack)-1] / curr
            }
            // reset current
            curr = 0
            // change the lastOp to new op
            lastOp = stringChar
        }
    }
    
   
    
    result := 0
    for len(stack) != 0 {
        result += stack[len(stack)-1]
        stack = stack[:len(stack)-1]
    }
    return result
}

func isDigit(s string) bool {
    return s >= "0" && s <= "9"
}