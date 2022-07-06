func subsets(nums []int) [][]int {
    result := [][]int{{}}
    for i := 0; i < len(nums); i++ {
        for _, eachList := range result {
            newL := make([]int, len(eachList))
            copy(newL, eachList)
            newL = append(newL, nums[i])
                    result = append(result, newL)

        }
    }
    return result
}