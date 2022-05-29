func twoSum(nums []int, target int) []int {
 
    if nums == nil || len(nums) == 0 {
        return nil
    }
    
    
    // go over each number
    // ask the following with the number in hand
    // what number do I need that when I add to current number
    // in hand, it will equal to target 
    // then the answer will be x = target-currentNum
    // I need the x such that when added with the currentNumber
    // i will have a sum of target
    
    // but since we need to do this with all number
    // we need a easy way to look up our x, so we will maintain
    // each number in a map. Because they can be the potential x
    // for someone..
    
    
//     idxMap := map[int]int{}
    
//     for i := 0; i < len(nums); i++ {
//         num := nums[i]
//         // what number do I need that when I add to current number
//         // in hand, it will equal to target 
//         x := target - num
        
//         // have we seen this x in our idx map before?
//         if idx, found := idxMap[x]; found {
//             return []int{idx, i}
//         }
        
//         // save the current number as it maybe a complement for some other
//         // number in the array down the line
//         idxMap[num] = i
//     }
    
//     return nil
    
    
    // follow up -- count the number of pairs possible...
    idxMap := map[int]int{}
    pairs := [][]int{}
    out := []int{}
    
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        // what number do I need that when I add to current number
        // in hand, it will equal to target 
        x := target - num
        
        // have we seen this x in our idx map before?
        if idx, found := idxMap[x]; found {
            pairs = append(pairs, []int{nums[idx], num})
            if len(out) == 0 {
                out = append(out, idx,i)
            }
        }
        
        // save the current number as it maybe a complement for some other
        // number in the array down the line
        idxMap[num] = i
    }
    // fmt.Println("All the pairs that add up to target: ", pairs)
    return out
}