func minCostClimbingStairs(cost []int) int {
    dp := make([]int, len(cost))
    dp[len(dp)-1] = cost[len(cost)-1]
    
    for i := len(cost)-2; i >= 0; i-- {
        oneStepCost := dp[i+1]
        twoStepCost := 0
        if i+2 < len(cost) {
            twoStepCost = dp[i+2]
        }
        currentCost := cost[i]
        
        if i == 0 {
            ifWeChoose0thIdx := min(oneStepCost,twoStepCost)+currentCost
            dp[i] =  min(oneStepCost,ifWeChoose0thIdx )
        } else {
            dp[i] = int(math.Min(float64(oneStepCost), float64(twoStepCost)))+currentCost
        }
    }
    return dp[0]
}

func min(x, y int) int{
    if x < y {return x}
    return y
}