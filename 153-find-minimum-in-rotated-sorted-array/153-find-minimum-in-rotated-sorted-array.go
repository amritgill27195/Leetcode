func findMin(nums []int) int {
    left := 0
    right := len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[left] <= nums[right] {
            return nums[left]
        }
        
        if (mid==0 || nums[mid] < nums[mid-1]) && (mid==len(nums)-1 || nums[mid] < nums[mid+1]) {
            return nums[mid]
        } else if nums[left] <= nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}