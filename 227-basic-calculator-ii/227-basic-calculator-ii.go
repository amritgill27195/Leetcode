func calculate(s string) int {
    op := "+"
    curr := 0
    st := []int{}
    for idx, char := range s {
        charString := string(char)
        charN, _ := strconv.Atoi(charString)
        
        if isDigit(charString) {
            curr = curr * 10 + charN
        }
        if (charString == "+" || charString == "-" || charString == "*" || charString == "/") || idx == len(s)-1 {
            if op == "+" {
                st = append(st, curr)
            } else if op == "-" {
                st = append(st, -curr)
            } else if op == "*" {
                st[len(st)-1] *= curr
            } else if op == "/" {
                st[len(st)-1] = st[len(st)-1]/curr
            }
            curr = 0
            op = charString
        }
    }
    
    total := 0
    for i := 0; i < len(st); i++ {
        total += st[i]
    }
    return total
}

func isDigit(s string) bool {
    return s >= "0" && s <= "9"
}