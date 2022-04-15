type TrieNode struct {
    isEnd bool
    childrens [26]*TrieNode
}

func insertWord(word string, root *TrieNode) {
    curr := root
    for i := 0; i < len(word); i++ {
        char := word[i]
        idx := char-'a'
        if curr.childrens[idx] == nil {
            curr.childrens[idx] = &TrieNode{
                childrens: [26]*TrieNode{},
            }
        }
        curr = curr.childrens[idx]
    }
    curr.isEnd = true

}

func longestWord(words []string) string {
    root := new(TrieNode)
    root.childrens = [26]*TrieNode{}
    root.isEnd = false
    
    for _, word := range words {
        insertWord(word, root)
    }
    maxWord := ""
    var backtrack func(r *TrieNode, pathBldr string)
    backtrack = func(r *TrieNode, pathBldr string) {
        // base
        if len(pathBldr) > len(maxWord) { 
            maxWord = pathBldr
        }

        
        // logic
        for i := 0; i <= 25; i++ {
            if r.childrens[i] != nil && r.childrens[i].isEnd {
                pathBldr += string('a'+i)
                backtrack(r.childrens[i], pathBldr)
                pathBldr = pathBldr[:len(pathBldr)-1]
            }
        }
    }
    backtrack(root, "")
    return maxWord
}