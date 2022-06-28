// top down from the back of the days array - bubbling the ans to idx 0
// func mincostTickets(days []int, costs []int) int {
    
//     travelDates := map[int]struct{}{}
//     for i := 0; i < len(days); i++ {
//         travelDates[days[i]] = struct{}{}
//     }
    
//     // bottom up dp
//     // each subproblem has 3 choices, make the best one for each ( i.e min cost )
//     dp := make([]int, days[len(days)-1])
//     n := len(dp)
//     oneDayPass := costs[0]
//     sevenDayPass := costs[1]
//     thirtyDayPass := costs[2]
//     dp[len(dp)-1] = min(oneDayPass,min(sevenDayPass,thirtyDayPass))

//     for i := len(dp)-2; i >= 0; i-- {
//         travelDate := i+1
//         _, traveling := travelDates[travelDate]
//         if !traveling { dp[i] = dp[i+1]; continue }
    
//         oneDayTotalCost := oneDayPass + dp[i+1]
//         sevenDayTotalCost := sevenDayPass
//         if i+7 < n { sevenDayTotalCost = sevenDayPass + dp[i+7] }
//         thirtyDayTotalCost := thirtyDayPass
//         if i+30 < n {thirtyDayTotalCost = thirtyDayPass + dp[i+30] }
//         dp[i] = min(oneDayTotalCost,min(sevenDayTotalCost,thirtyDayTotalCost))
//     }
//     return dp[0]
// }

func mincostTickets(days []int, costs []int) int {
    dp := make([]int, days[len(days)-1]+1)
    
    oneDayPass := costs[0]; sevenDayPass := costs[1]; thirtyDayPass := costs[2]
    dp[days[0]] = min(oneDayPass, min(sevenDayPass, thirtyDayPass))
    
    dates := map[int]struct{}{}
    for i := 0; i < len(days); i++ {
        dates[days[i]] = struct{}{}
    }
    
    
    for i := days[0]+1; i < len(dp); i++ {
        _, ok := dates[i]
        if !ok {
            dp[i] = dp[i-1]
            continue
        }
        
        oneTotal := oneDayPass + dp[i-1]
        sevenTotal := sevenDayPass
        if i-7 >= 0 {
            sevenTotal +=  dp[i-7]
        }
        thirtyTotal := thirtyDayPass
        if i-30 >= 0 {
            thirtyTotal += dp[i-30]
        }
        
        dp[i] = min(oneTotal, min(thirtyTotal,sevenTotal))
    }
    return dp[len(dp)-1]
}

func min(x, y int) int{
    if x < y {return x}
    return y
}