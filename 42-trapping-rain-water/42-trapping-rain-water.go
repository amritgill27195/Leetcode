/*
    approach: 2 pass using 2 pointers ( slow and fast ptrs )
    - We want to position our 2 pointers such that we know we can trap some water
    - fast pointer will always move
    - if height at fast pointer is smaller than height at left
    - this means we can trap some water between left and right - so we will take whatever we can trap here
        - trap += height[left]-height[fast]
    - when height at fast is not less than height at slow, this means we are done trapping as much water as we can , so add our trapped value to final result
    - move slow to fast and reset trap value
    - always move fast pointer regardless of which path we take
    - but this does not work in a single pass everytime... because we could have a slow pointer pointing the tallest bar in the middle of the array
    - and all bars to right are always less than slow height ( i.e height[fast] is always less than height[slow] because slow is at the tallest point )
    - Therefore another pass from the back is needed to check if we can trap water the other until this peek ( if there is a peek value slow paused at )
    - so loop back until peek idx and repeat the same checks
   
   time: o(2n) - but 2 is constant
   space: o(1)
   
*/ 
func trap(height []int) int {
    n := len(height)
    result := 0
    trap := 0
    
    slow := 0 // acting as the left wall
    fast := slow+1
    
    for fast < n {
        if height[slow] > height[fast] {
            trap += height[slow]-height[fast]
        } else {
            result += trap
            trap = 0
            slow = fast
        }
        fast++
    }
    
    peek := slow
    slow = n-1 // acting as the right wall now
    fast = slow - 1
    trap = 0
    for fast >= peek {
        if height[slow] > height[fast] {
            trap += height[slow]-height[fast]
        } else {
            result += trap
            trap = 0
            slow = fast
        }
        fast--
    }
    
    return result
}