/*
    Sorted array, then we should consider binary search and or two pointers
    
    approach: hashmap and complement search
    - for each number in nums array,
    - diff = target-number
    - if this diff exists ( i.e the complement ) in our hashmap, then return this and hashmap value as idx
    - otherwise add this number and its index in hashmap
    
    time: o(n)
    spaceL o(n)

*/

// func twoSum(numbers []int, target int) []int {
//     m := map[int]int{}
//     for i := 0; i < len(numbers); i++ {
//         diff := target - numbers[i]
//         if idx, ok := m[diff]; ok {
//             return []int{idx+1, i+1}
//         }
//         m[numbers[i]]=i
//     }
//     return []int{-1,-1}
// }


/*
    Sorted array, then we should consider binary search and or two pointers
    
    approach: complement search using binary search
    - for each num in numbers
    - calc the diff, target - num ( complement )
    - search for that diff using binary search
    
    time: o(nlogn)
    spaceL o(1)

*/

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        diff := target - nums[i]
        left := i+1
        right := len(nums)-1
        for left <= right {
            mid := left + (right-left)/2
            if nums[mid] == diff {
                return []int{i+1, mid+1}
            } else if nums[mid] > diff {
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return []int{-1,-1}
}