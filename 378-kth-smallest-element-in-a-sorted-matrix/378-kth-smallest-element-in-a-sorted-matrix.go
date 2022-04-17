
/*
    approach: brute force
    - Flatten the entire matrix in a 1D array
    - Sort it
    - Return the k+1 idx element
    Time: o(mn) + o(nlogn)
    Space: o(mn)
    
    
    
*/


// brute force
func kthSmallest(matrix [][]int, k int) int {
    flat := []int{}
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            flat = append(flat, matrix[i][j])
        }
    }
    sort.Ints(flat)
    return flat[k-1]
}