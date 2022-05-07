func search(nums []int, target int) int {
    left := 0
    right := len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] >= nums[left] {
            // left sorted
            if target >= nums[left] && target < nums[mid] {
                right = mid-1
            } else {
                left = mid+1
            }
            
        } else {
            // right sorted
            if target <= nums[right] && target > nums[mid] {
                left = mid+1
            } else {
                right = mid-1
            }
        }
    }
    return -1
}