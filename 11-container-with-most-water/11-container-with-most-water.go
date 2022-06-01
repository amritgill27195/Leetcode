func maxArea(height []int) int {
    max := 0
    left := 0
    right := len(height)-1
    
    for left < right {
        h := int(math.Min(float64(height[left]), float64(height[right])))
        w := right-left
        area := h*w
        if area > max {
            max = area
        }
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return max
}   