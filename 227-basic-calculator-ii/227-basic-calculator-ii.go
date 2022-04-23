func calculate(s string) int {
    calc := 0
    current := 0
    tail := 0
    op := "+"
    
    for i := 0; i < len(s); i++ {
        charString := string(s[i])
        if unicode.IsDigit(rune(s[i])) {
            n, _ := strconv.Atoi(charString)
            current = current * 10 + n
            
        }
        if (!unicode.IsDigit(rune(s[i])) && charString != " ") || i == len(s)-1 {
            if op == "+" {
                calc += current
                tail = current
            } else if op == "-" {
                calc -= current
                tail = -current
            } else if op == "*" {
                calc = calc - tail + (tail*current)
                tail = tail * current
            } else if op == "/" {
                calc = calc - tail + (tail/current)
                tail = tail / current
            }
            current = 0
            op = charString
        }
    }
    return calc
}