func subsets(nums []int) [][]int {
    
    result := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, el := range result {
            newL := make([]int, len(el))
            copy(newL, el)
            newL = append(newL, nums[i])
            result = append(result, newL)
        }
    }
    
    return result
}