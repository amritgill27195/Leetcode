/*
    approach: brute force
    - nested for loops
    - check if jth element is greater than ith element
    - if it is, grab the distance (j-i) and save it in out[i]
    time: o(n^2)
    space: o(1)
    
    ( also this is essentially, find the next greater element on right but instead of value, save the distance between two points )
    approach: monotonic increasing stack
    - Loop over the list
    - if the stack is not empty, and this item resolves previous its immediate previous neighbour ( top of the stack )
        - resolve in this context means, its warmer than previous day
        - Then pop the top of the stack 
        - Calculate distance between top and current item ( currentIdx - idxPoppedFromTopOfStack )
        - Save the distance in the output array at idxPoppedFromTopOfStack
    - Otherwise save the currentIdx in the stack
    
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