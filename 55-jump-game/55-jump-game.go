// DFS + Memo -- Result still TLE
// func canJump(nums []int) bool {
//     memo := map[int]bool{}
//     var dfs func(idx int) bool
//     dfs = func(idx int) bool {
//         // base
//         if idx >= len(nums)-1 {return true}
        
//         // logic
//         for i := nums[idx]; i >= 1; i-- {
//             newIdx := idx+i
//             if ok, val := memo[newIdx]; ok && val == true{
//                 return true
//             } else {
//                 res := dfs(newIdx)
//                 if res {return true}
//                 memo[newIdx] = res
//             }
//         }
//         return false
//     }
//     return dfs(0)
// }

func canJump(nums []int) bool { 
    dest := len(nums)-1
    // fmt.Println("dest: ", dest)
    for i := len(nums)-2; i >= 0; i-- {
        numJumpsAvail := nums[i]
        if i+numJumpsAvail >= dest {
            dest = i
        }
    }
    return dest == 0
}
