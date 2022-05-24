/*
    approach: slow and fast pointers
    - fast pointer becomes out iterating pointer
    - using count ( starting with 1 since we have already seen 0th idx - implied )
    - if fast ptr  == fast-1 , count++
    - count is keeping track of number of dupes by comparing current fast with previous fast
    - if count <= k
        - Then we will move slow by 1
        - And copy the value from fast to slow pointer
    - otherwise if current fast val != prev fast value, then reset count to 1
    - finally return slow+1
    
*/
func removeDuplicates(nums []int) int {
    slow := 1
    fast := 1
    count := 1
    k := 2
    n := len(nums)
    for fast < n {
        if nums[fast] == nums[fast-1] {
            count++
        } else {
            count = 1
        }
        if count <= k {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return slow
}