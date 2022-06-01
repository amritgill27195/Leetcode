func trap(height []int) int {
    n := len(height)
    result := 0
    slow := 0
    fast := slow+1
    trap := 0
    for fast < n {
        if height[fast] < height[slow] {
            trap += height[slow]-height[fast]
        } else {
            result += trap
            slow = fast
            trap = 0
        }
        fast++
    }
    peek := slow
    slow = n-1
    fast = slow-1
    trap = 0
    for fast >= peek {
        if height[fast] < height[slow] {
            trap += height[slow] - height[fast]
        } else {
            result += trap
            slow = fast
            trap = 0
        }
        fast--
    }
    return result
}