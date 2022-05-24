/*
    approach: slow and fast pointers
    - slow and fast pointer both starting at idx 1
    - count goes up by 1 if fast == fast-1
    - count resets to 1 if fast != fast-1
    - if count <= k, this tells slow pointer that it needs to move and that slow pointer has not seen k dupes to be allowed yet
        - replace value at slow with fast and slow++
    
    time: o(n)
    space: o(1)
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