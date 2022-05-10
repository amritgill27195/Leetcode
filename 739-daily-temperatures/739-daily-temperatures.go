/*
    approach: brute force
    - nested for loops
    - check if jth element is greater than ith element
    - if it is, grab the distance (j-i) and save it in out[i]
    time: o(n^2)
    space: o(1)
    
    ( also this is essentially, find the next greater element on right but instead of value, save the distance between two points )
    approach: monotonic increasing stack
    time: o(n)
    space: o(n)

*/

func dailyTemperatures(temperatures []int) []int {
    out := make([]int, len(temperatures))
    stack := []int{}
    for i := 0; i < len(temperatures); i++ {
        for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            out[top] = i-top
        }
        stack = append(stack, i)
    }
    return out
}