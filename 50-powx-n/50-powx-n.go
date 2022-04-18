/*
    Why is this a question.... I will never ever implement my math.Pow... 
    Interviews are getting out of control... these days but this is the game we have to play...
    so it is what it is.... sigh
    
    approach: brute force 
    - loop n times and multiply x...
    - if n < 0 , then x = 1/x and n = -n
    - for loop 0 to n(exlcuding n and multiply x each time to a result var)
    - return result
    time: o(n) - we loop n times
    space: o(1)
    - Outcome: TLE
    
    approach: logN recursive
    
    
    
*/
// brute force
// func myPow(x float64, n int) float64 {
//     if n < 0 {
//         x = 1/x
//         n *= -1
//     }
//     result := 1.0
//     for i := 0; i < n; i++ {
//         result *= x
//     }
//     return result
// }

func myPow(x float64, n int) float64 {
    // base
    if n == 0 {return 1.00}
    // logic
    result := myPow(x, n/2)
    if n % 2 == 0{
        return result * result
    } else {
        if n < 0 {
            return result * result * 1/x
        } else {
            return result * result * x
        }
    }
}