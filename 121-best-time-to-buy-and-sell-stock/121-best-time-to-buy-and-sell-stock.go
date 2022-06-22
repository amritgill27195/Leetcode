func maxProfit(prices []int) int {
    profit := 0
    buy := prices[0]
    for i := 1; i < len(prices); i++ {
        sell := prices[i]
        cp := sell-buy
        if cp > profit {profit = cp}
        if sell < buy {buy = sell}
    }
    
    return profit
}