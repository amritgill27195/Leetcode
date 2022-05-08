func maxProfit(prices []int) int {
    max := 0
    buy := prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i]-buy > max {
            max = prices[i]-buy
        }
        if prices[i] < buy {
            buy = prices[i]
        }
    }
    return max
}