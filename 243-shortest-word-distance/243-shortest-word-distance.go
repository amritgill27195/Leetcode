/*
    
*/

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
    idxMap := map[string][]int{}
    for i := 0; i < len(wordsDict); i++ {
        idxMap[wordsDict[i]] = append(idxMap[wordsDict[i]], i)
    }
    
    p1 := 0
    p1List := idxMap[word1]
    
    p2 := 0
    p2List := idxMap[word2]
    
    minDist := math.MaxInt64
    for p1 < len(p1List) && p2 < len(p2List) {
        p1Idx := p1List[p1]
        p2Idx := p2List[p2]
        dist := abs(p1Idx-p2Idx)
        if dist < minDist {
            minDist = dist
        }
        if p1Idx < p2Idx {
            p1++
        } else {
            p2++
        }
    }
    return minDist
}

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}