func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    result := 0
    for i := 0; i < len(nums); i+=2 {
        min := int(math.Min(float64(nums[i]), float64(nums[i+1])))
        result += min
    }
    return result
}