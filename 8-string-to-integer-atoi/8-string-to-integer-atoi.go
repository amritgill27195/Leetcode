// time: o(n)
// space: o(1)
func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if s == "" {
        return 0
    }
    first := s[0]
    if !(first == '-' || first == '+') && !isDigit(first) {
        return 0
    }
    isNegative := false
    if first == '-' {isNegative=true}
    n := 0
    limit := math.MaxInt32/10
    for i := 0; i < len(s); i++ {
        char := s[i]

        if isDigit(char) {
            charInt := int(char-'0')
            if !isNegative {
                /*
                    why these 2 checks as separate checks?
                    - first, maxInt32 is 2147483647
                    - keep in mind we are "APPENDING" numbers to our n
                    - we drop the last number of maxInt32 (maxInt32/10) and we are left with 214748364
                        - why?
                        - because we have 2 cases
                        - if our n == limit(214748364) then, the only choice of appending another number must be within 0 to 7 ( inclusive )
                            - so explicitly check if the new int being appending is within this range, if not, return maxInt32
                            - but then, why are we comparing with 8 when the number being formed in negative?
                            - because minInt32 (when we are forming negative number), minInt32 limit is 2147483648 instead of 2147483647
                            - therefore...
                        - We are appending, not adding...
                        - the second case is when our already past limit ( n > limit(214748364) )
                            - i.e n could be 214748365, 214748366, 214748367 for example
                            - then we cannot proceed to add append another number right? 
                        - therefore in this case, as soon as we are past the limit, regardless if the new number being appending is within our range, we cannot add it because
                        - it will be past int32.
                    
                        
                */
                if n == limit && charInt >= 7 {return math.MaxInt32}
                if n > limit {return math.MaxInt32}
            } else {
                if n == limit && charInt >= 8 {return math.MinInt32}
                if n > limit {return math.MinInt32}
            }
            n = n * 10 + charInt
        } else if i != 0 {
            break
        }
    }
    if isNegative {
        n *= -1
    }
    return n
    
}

func isDigit(n byte) bool {
    return n >= '0' && n <= '9'
}