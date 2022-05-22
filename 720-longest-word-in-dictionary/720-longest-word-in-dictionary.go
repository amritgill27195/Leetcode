type trieNode struct {
    childrens [26]*trieNode
    isEnd bool
}

func insert(root *trieNode, word string) {
    if root == nil {return}
    curr := root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childrens[charIdx] == nil {
            curr.childrens[charIdx] = &trieNode{childrens: [26]*trieNode{}}
        }
        curr = curr.childrens[charIdx]
    }
    curr.isEnd = true
}


func longestWord(words []string) string {
    
    root := &trieNode{childrens: [26]*trieNode{}}
    for _, word := range words {
        insert(root, word)
    }
    
    result := ""
    var backtrack func(path string, r *trieNode)
    backtrack = func(path string, r *trieNode) {
        // base
        if len(path) > len(result) {
            result = path
        }
        if r == nil {return}
        
        // logic
        curr := r
        for i := 0; i < 26; i++ {
            if curr.childrens[i] != nil && curr.childrens[i].isEnd {
                // action
                path += string(i+'a')
                // recurse
                backtrack(path,curr.childrens[i])
                // backtrack
                path = string(path[:len(path)-1])
            }
        }
    }
    backtrack("", root)
    return result
    
}