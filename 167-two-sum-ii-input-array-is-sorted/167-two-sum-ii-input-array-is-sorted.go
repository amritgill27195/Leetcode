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

// func twoSum(nums []int, target int) []int {
//     for i := 0; i < len(nums); i++ {
//         diff := target - nums[i]
//         left := i+1
//         right := len(nums)-1
//         for left <= right {
//             mid := left + (right-left)/2
//             if nums[mid] == diff {
//                 return []int{i+1, mid+1}
//             } else if nums[mid] > diff {
//                 right = mid-1
//             } else {
//                 left = mid+1
//             }
//         }
//     }
//     return []int{-1,-1}
// }



/*
    Sorted array, then we should consider binary search and or two pointers
    
    approach: two pointers
    - take a left and right pointer ( start and end of nums array )
    - if their sum == target, return those indicies
    - otherwise if sum > target 
        - move right pointer back as moving this pointer back will result into smaller sum
        - thereby increasing the chances of being equal to target
    - else move left pointer forward
        - because if their sum is smaller than the target
        - moving right pointer back in a sorted array will only make the sum smaller
        - moving left pointer forward will help in this case and increase chances of getting equal to target
    
    
    time: o(n)
    spaceL o(1)

*/

func twoSum(nums []int, target int) []int {
    left := 0
    right := len(nums)-1
    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {
            return []int{left+1, right+1}
        } else if sum > target {
            right--
        } else {
            left++
        }
    }
    return nil
}