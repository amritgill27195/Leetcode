// func containsDuplicate(nums []int) bool {
    
    
//     // hashmap
//     set := map[int]struct{}{}
//     for i := 0; i < len(nums); i++ {
//         if _, seen := set[nums[i]]; seen {
//             return true
//         }
//         set[nums[i]] =struct{}{}
//     }
    
//     return false
// }

// sorting
func containsDuplicate(nums []int) bool {
    sort.Ints(nums)
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}