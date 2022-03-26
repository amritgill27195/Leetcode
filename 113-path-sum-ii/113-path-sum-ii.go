// func pathSum(root *TreeNode, targetSum int) [][]int {
// 	var result [][]int
//     inorderDfs(root,targetSum,nil, &result)
// 	return result
// }
// func inorderDfs(root *TreeNode, targetSum int, paths []int, result *[][]int) {
// 	if root == nil {
// 		return
// 	}
//         fmt.Println("root: ", root.Val)
//     fmt.Printf("before adding root -- paths: %v and slice mem address: %p len: %d and cap: %d\n", paths, paths, len(paths), cap(paths) )

// 	targetSum -= root.Val
//     paths = append(paths, root.Val)

//     fmt.Printf("after adding root -- paths: %v and slice mem address: %p len: %d and cap: %d\n", paths, paths, len(paths), cap(paths) )

//     inorderDfs(root.Left, targetSum, paths, result)
    
    
//     if targetSum == 0 && root.Left == nil && root.Right == nil {
//         newPath := make([]int, len(paths))
//         copy(newPath, paths)
// 		*result = append(*result, newPath)
// 	}
    
    
//     inorderDfs(root.Right, targetSum, paths, result)


// }


func pathSum(root *TreeNode, sum int) [][]int {
    var paths [][]int

    if root == nil {
        return paths
    }

    if root.Left == nil && root.Right == nil {
        if sum == root.Val {
            return append(paths, []int{ root.Val })
        }

        return paths
    }

    for _, path := range pathSum(root.Left, sum - root.Val) {
        paths = append(paths, append([]int{ root.Val}, path... ))
    }

    for _, path := range pathSum(root.Right, sum - root.Val) {
        paths = append(paths, append([]int{ root.Val}, path... ))
    }

    return paths
}