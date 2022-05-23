func merge(nums1 []int, m int, nums2 []int, n int)  {
    back := len(nums1)-1
    n1 := m-1
    n2 := n-1
    
    for back >= 0 {
        if n1 >= 0 && n2 >= 0 {
            n1Val := nums1[n1]
            n2Val := nums2[n2]

            if n1Val > n2Val {
                nums1[back] = n1Val
                n1--
            } else {
                nums1[back] = n2Val
                n2--
            }
        } else if n1 >= 0 {
            nums1[back] = nums1[n1]
            n1--
        } else if n2 >= 0 {
            nums1[back] = nums2[n2]
            n2--
        }
        back--
    }

}