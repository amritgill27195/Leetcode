func mincostTickets(days []int, costs []int) int {
    
    travelDates := map[int]struct{}{}
    for i := 0; i < len(days); i++ {
        travelDates[days[i]] = struct{}{}
    }
    // fmt.Println(travelDates)
    
    // bottom up dp
    // each subproblem has 3 choices, make the best one for each ( i.e min cost )
    dp := make([]int, days[len(days)-1])
    n := len(dp)
    oneDayPass := costs[0]
    sevenDayPass := costs[1]
    thirtyDayPass := costs[2]
    dp[len(dp)-1] = min(oneDayPass,min(sevenDayPass,thirtyDayPass))

    
    for i := len(dp)-2; i >= 0; i-- {
        
        travelDate := i+1
        _, traveling := travelDates[travelDate]
        if traveling {
            
            oneDayTotalCost := oneDayPass + dp[i+1]
            
            sevenDayTotalCost := 0
            sevenDaysFromNowIdx := i+7-1
            if sevenDaysFromNowIdx >= n-1 {
                sevenDayTotalCost = sevenDayPass
            } else {
                sevenDayTotalCost = sevenDayPass + dp[sevenDaysFromNowIdx+1]
            }
            
            thirtyDayTotalCost := 0
            thirtyDaysFromNowIdx := i+30-1
            if thirtyDaysFromNowIdx >= n-1 {
                thirtyDayTotalCost = thirtyDayPass
            } else{
                thirtyDayTotalCost = thirtyDayPass + dp[thirtyDaysFromNowIdx+1]
            }
            
            // fmt.Println(oneDayTotalCost,sevenDayTotalCost, thirtyDayTotalCost)
            // fmt.Println("------------------------------------------------")
            
            dp[i] = min(oneDayTotalCost,min(sevenDayTotalCost,thirtyDayTotalCost))
            
            
        } else {
            dp[i] = dp[i+1]
        }
    }
    // fmt.Println(dp)
    return dp[0]
}

func min(x, y int) int{
    if x < y {return x}
    return y
}