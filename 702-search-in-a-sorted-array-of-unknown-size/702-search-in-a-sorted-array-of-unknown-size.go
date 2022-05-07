/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
    left := 0
    right := left+1
    for reader.get(right) < target {
        left = right
        right = right * 2
    }
    
    for left <= right {
        mid := left + (right-left)/2
        midVal := reader.get(mid)
        if midVal == target {
            return mid
        } else if midVal > target {
            right = mid-1
        }  else {
            left = mid+1
        }
    }
    return -1
}