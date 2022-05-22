func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {
        return intersect(nums2, nums1)
    }
    
    sort.Ints(nums1)
    sort.Ints(nums2)
    
    out := []int{}
    left := 0
    
    for _, ele := range nums1 {
        idx := searchLeftestIdx(left, ele, nums2)
        if idx != -1 {
            left = idx+1
            out = append(out, ele)
        }
    }

    
    return out
}

func searchLeftestIdx(left, target int, nums []int) int {
    right := len(nums)-1
    idx := -1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            if nums[mid] == target {
                idx = mid
            }
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}