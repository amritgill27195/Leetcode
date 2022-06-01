func twoSum(numbers []int, target int) []int {
    for i := 0; i < len(numbers); i++ {
        num := numbers[i]
        complement := target-num
        left := i+1
        right := len(numbers)-1
        for left <= right {
            mid := left + (right-left)/2
            if numbers[mid] == complement {
                return []int{i+1, mid+1}
            } else if numbers[mid] > complement {
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return []int{-1,-1}
}