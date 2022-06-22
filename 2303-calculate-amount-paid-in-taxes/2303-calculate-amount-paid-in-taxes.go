func calculateTax(brackets [][]int, income int) float64 {
    tax := 0.0
    i := 0
    for income > 0 {
        taxRate := float64(brackets[i][1])/100.00
        taxBracket := brackets[i][0]
        if i != 0 {
            taxBracket = abs(brackets[i][0]-brackets[i-1][0])
        }
        min := int(math.Min(float64(income), float64(taxBracket)))
        tax += (float64(min) * float64(taxRate))
        i++
        income -= min
    }
    
    return tax
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}