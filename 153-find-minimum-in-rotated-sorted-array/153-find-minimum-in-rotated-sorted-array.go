func findMin(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    
    for left <= right {
        if nums[left] <= nums[right] {
            return nums[left]
        }
        mid := left + (right-left)/2
        if (mid == 0 || nums[mid] < nums[mid-1]) && (mid == n-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        } else if nums[mid] >= nums[left] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}