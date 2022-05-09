
/*
    We can do as many transactions as possible...
    This is when we will compare buy on i and sell on i+1 price
    if there is any profit, add to total profit..
    
    Question becomes, if you sold it on a day, can you buy it again after selling it?
    The answer is not very clear but it is true, you can buy the stock after selling it

    time: o(n)
    space: o(1)

*/
func maxProfit(prices []int) int {
    totalProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            totalProfit +=  prices[i] - prices[i-1]
        }
    }
    
    return totalProfit
}