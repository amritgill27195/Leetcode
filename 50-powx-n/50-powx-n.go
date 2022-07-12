// func myPow(x float64, n int) float64 {
//     if n < 0 {
//         n *= -1
//         x = 1/x
//     }
//     var dfs func(base float64, exp int) float64
//     dfs = func(base float64, exp int) float64 {
//         // base
//         if exp == 1 {
//             return base
//         }
//         if exp == 0 {return 1}
        
//         // logic
//         if exp % 2 == 0 {
//             return dfs(base, exp/2) * dfs(base, exp/2)   
//         }
//         return base * dfs(base, exp/2) * dfs(base, exp/2)
//     }
//     return dfs(x, n)
    
// }

func myPow(x float64, n int) float64 {
    // if n is negative 
    // make it positive and make x = 1/x
    if n < 0 {
        n *= -1
        x = 1/x
    }
    
    var dfs func(base float64, exp int) float64
    dfs = func(base float64, exp int) float64 {
        // base
        if exp == 0 {
            return 1
        }
        
        // logic
        result := dfs(base, exp/2)
        if exp % 2 == 0 {
            return result * result
        } else {
            return result * result * base
        }
    }
    return dfs(x, n)
}