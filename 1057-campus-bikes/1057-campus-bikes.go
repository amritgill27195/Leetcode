func assignBikes(workers [][]int, bikes [][]int) []int {
    distMap := map[int][][]int{}
    min := math.MaxInt64
    max := math.MinInt64
    for wId, worker := range workers {
        for bId, bike := range bikes {
            dist := abs(worker[1]-bike[1]) + abs(worker[0]-bike[0])
            distMap[dist] = append(distMap[dist], []int{wId, bId})
            if dist < min {min = dist}
            if dist > max {max = dist}
        }
    }
    workersToBikeId := make([]int, len(workers))
    workersSet := map[int]bool{}
    bikesSet := map[int]bool{}
    for i := min; i <= max; i++ {
        wbPairs := distMap[i]
        for _, wbPair := range wbPairs {
            wId := wbPair[0]
            bId := wbPair[1]
            
            _, wIdInSet := workersSet[wId]
            _, bIdInSet := bikesSet[bId]
            if !wIdInSet && !bIdInSet {
                workersToBikeId[wId] = bId
                workersSet[wId] = true
                bikesSet[bId] = true
            }
        }
    }
    return workersToBikeId
}

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}