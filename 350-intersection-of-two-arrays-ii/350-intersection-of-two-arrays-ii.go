/*
    We need to find the common elements between nums1 and nums2
    
    approach: freqMap
    - take smaller array and turn it into freqMap
    - Then loop over the larger array, and check if the each element exists in freqMap
    - if it does, then add this element to our result array 
    - decrement this element's frequency in freqMap by 1
    - if the frequency of this element reaches 0, then delete this element from map so we do not accidentally pick it again if this element exists more than N times in larger array compared to smaller array
    - Finally return the result array which will only have the common elements 
    
    m = larger array
    n = smaller array
    
    time: o(n) + o(m)
    space: o(n)
*/

func intersect(nums1 []int, nums2 []int) []int {
    
    // to make sure nums1 is always the larger array
    if len(nums1) < len(nums2) {
        return intersect(nums2, nums1)
    }
    
    
    freqMap := map[int]int{}
    for _, num := range nums2 {
        freqMap[num]++
    }
    
    result := []int{}
    for i := 0; i < len(nums1); i++ {
        val, ok := freqMap[nums1[i]]
        if ok && val != 0 { // no need to delete, we can just check its frequency
            result = append(result, nums1[i])
            freqMap[nums1[i]]--
        }
    }
    return result
    
}