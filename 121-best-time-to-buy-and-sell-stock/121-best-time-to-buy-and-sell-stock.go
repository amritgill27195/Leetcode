func maxProfit(prices []int) int {
    buyIdx := 0
    buyDay := -1
    sellDay := -1
    max := 0
    for i := 1; i < len(prices); i++ {
        buy := prices[buyIdx]
        sell := prices[i]
        
        profit := sell - buy
        if profit > 0 {
            if profit > max {
                max = profit
                buyDay = buyIdx+1
                sellDay = i+1
            }
        }
        if sell < prices[buyIdx] {
            buyIdx = i
        }
    }
    fmt.Println("Max profit would be made if we bought it on day", buyDay," and sold it on day ", sellDay)
    return max
    
}