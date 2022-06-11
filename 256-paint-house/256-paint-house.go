func minCost(costs [][]int) int {
    m := len(costs)
    for i := m-2; i >= 0; i-- {
        costs[i][0] += int(math.Min(float64(costs[i+1][1]),float64(costs[i+1][2])))
        costs[i][1] += int(math.Min(float64(costs[i+1][0]),float64(costs[i+1][2])))
        costs[i][2] += int(math.Min(float64(costs[i+1][0]),float64(costs[i+1][1])))
    }
    
    return int(math.Min(float64(costs[0][0]), math.Min(float64(costs[0][1]), float64(costs[0][2]))))
}