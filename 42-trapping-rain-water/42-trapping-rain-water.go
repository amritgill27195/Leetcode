/*
    approach : 2 pointers 
    - left and right ptrs ( 0 and end )
    - We will have a leftMax val acting as the biggest left wall seen so far
    - The leftMax will be used as a left wall for the left ptr
    - i.e left ptr becomes the right wall and leftMax becomes the left wall
    - Then to check how much we can trap between leftMax and left 
    - We can simply do leftMax - left and add the answer to result
    - The inverse applies on the other side
    - We will have a rightMax which will tell us the biggest wall right ptr has seen so far
    - rightMax acts as the right wall for the right ptr and that means right ptr becomes the left wall and rightMax becomes the right wall
    - Then to check how much water is trapped between rightMax and right
    - rightMax - right will give us that ans and we simply add this to the result variable.
    
    - Which pointer ( left vs right ) moves when?
    - We move the smaller value
    - if both have the same value, move either one
    
    time: o(n)
    space: o(1)
*/

func trap(height []int) int {
    left := 0
    right := len(height)-1
    leftMax := height[left]
    rightMax := height[right]
    totalTrapped := 0
    for left < right {
        if height[left] < height[right] {
            left++
            if height[left] > leftMax {
                leftMax = height[left]
            }
            totalTrapped += leftMax - height[left]
        } else {
            right--
            if height[right] > rightMax {
                rightMax = height[right]
            }
            totalTrapped += rightMax - height[right]
        }
    }
    return totalTrapped
}


/*
    approach: slow and fast pointers : 2 pass
    
*/
// func trap(height []int) int {
    
//     slow := 0
//     fast := 1
//     n := len(height)
//     trap := 0
//     result := 0
    
//     for fast < n {
//         if height[fast] < height[slow] {
//             trap += height[slow]-height[fast]
//         } else {
//             result += trap
//             trap = 0
//             slow = fast
//         }
//         fast++
//     }
//     peak := slow
//     slow = n-1
//     fast = slow-1
//     trap = 0
//     for fast >= peak {
//         if height[fast] < height[slow] {
//             trap += height[slow]-height[fast]
//         } else {
//             result += trap
//             trap = 0
//             slow = fast
//         }
//         fast--
//     }
//     return result
    
// }