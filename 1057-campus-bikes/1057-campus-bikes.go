func assignBikes(workers [][]int, bikes [][]int) []int {
    minDist := math.MaxInt64
    maxDist := math.MinInt64
    distMap := map[int][][]int{}
    
    for wrI, wr := range workers {
        wrX := wr[0]
        wrY := wr[1]
        for bkI, bk := range bikes {
            bkX := bk[0]
            bkY := bk[1]
            dist := abs(bkY-wrY) + abs(bkX-wrX)
            if dist < minDist {minDist = dist}
            if dist > maxDist {maxDist = dist}
            
            distMap[dist] = append(distMap[dist], []int{wrI, bkI})
        }
    }
    
    workerToBikes := make([]int, len(workers))
    workersAssigned := make([]bool, len(workers))
    bikesOccupied := make([]bool, len(bikes))
    
    count := len(workers)
    for i := minDist; i <= maxDist; i++ {
        wbPairs := distMap[i]
        for _, wbPair := range wbPairs {
            wr := wbPair[0]
            bk := wbPair[1]
            if !workersAssigned[wr] && !bikesOccupied[bk] {
                workersAssigned[wr] = true
                bikesOccupied[bk] = true
                workerToBikes[wr] = bk
                count--
            }
            if count == 0 {return workerToBikes}
        }
    }
    return workerToBikes
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}