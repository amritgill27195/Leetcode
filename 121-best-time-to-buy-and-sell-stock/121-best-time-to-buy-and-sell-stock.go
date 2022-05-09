
/*
    since we have to return the largest possible profit,
    We want to buy it on the cheapest day possible and sell it on the day that gives us most profit
    
    We will start by implying that cheapest day is at idx 0
    so we will buy at price[0]
    
    Then loop over prices from 1, and check the profit if we were to sell on for ith price.
    (price[i]-buy) -- if this profit > max - then update max to the new number.
    
    However, because we implied that at 0th idx, the price is the cheapest, we need to also keep on checking
    if the current ith price is cheaper/smaller than the price we originally bought it for.
    If it is, save this as our new buy price and compare with rest of the selling prices down the line.
    
    Return the max at the end
    
    time: o(n)
    space: o(1)

*/
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