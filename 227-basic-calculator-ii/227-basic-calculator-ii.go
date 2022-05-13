func calculate(s string) int {
    op := "+"
    calc := 0
    curr := 0
    tail := 0
    for idx, char := range s {
        charString := string(char)
        charN, _ := strconv.Atoi(charString)
        
        if isDigit(charString) {
            curr = curr * 10 + charN
        }
        if (charString == "+" || charString == "-" || charString == "*" || charString == "/") || idx == len(s)-1 {
            if op == "+" {
                calc += curr
                tail = curr
            } else if op == "-" {
                calc -= curr
                tail = -curr
            } else if op == "*" {
                calc = calc - tail + (tail*curr)
                tail = tail * curr
            } else if op == "/" {
                calc = calc - tail + (tail/curr)
                tail = tail / curr
            }
            curr = 0
            op = charString
        }
    }
    
return calc
}

func isDigit(s string) bool {
    return s >= "0" && s <= "9"
}