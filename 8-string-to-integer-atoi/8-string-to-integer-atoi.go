// time: o(n)
// space: o(1)
func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if s == "" {
        return 0
    }
    first := s[0]
    if (!(first == '-' || first == '+')) && !isDigit(first) {
        return 0
    }
    n := 0
    isNegative := false
    if first == '-' { isNegative = true }
    limit := math.MaxInt32/10
    for i := 0; i < len(s); i++ {
        c := s[i]
        cInt := int(c-'0')
        if isDigit(c) {
            if !isNegative {
                if n > limit {return math.MaxInt32}
                if n == limit && cInt >= 7 {return math.MaxInt32}
            } else {
                if n > limit {return math.MinInt32}
                if n == limit && cInt >= 8 {return math.MinInt32}
            }
            n = n * 10 + cInt
        } else if i != 0 {
            break 
        }
    }
    if isNegative {
        n *= -1
    }
    return n
}
func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}