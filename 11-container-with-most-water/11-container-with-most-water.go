func maxArea(height []int) int {
    left := 0
    right := len(height)-1
    maxArea := 0
    for left < right {
        leftHeight := height[left]
        rightHeight := height[right]
        height := int(math.Min(float64(leftHeight), float64(rightHeight)))
        width := right-left
        if height*width > maxArea {
            maxArea = height*width
        }
        if leftHeight < rightHeight {
            left++
        } else {
            right--
        }
    }
    return maxArea
}