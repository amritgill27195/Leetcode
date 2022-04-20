type trieNode struct {
    isEnd bool
    childrens [26]*trieNode
}

func insert(root *trieNode, word string) {
    current := root
    for _, char := range word {
        if current.childrens[char-'a'] == nil {
            current.childrens[char-'a'] = &trieNode{childrens: [26]*trieNode{}}
        }
        current = current.childrens[char-'a']
    }
    current.isEnd = true
}


func search(root *trieNode, word string) (bool, string) {
    current := root
    out := new(strings.Builder)
    for _, char := range word {
        if current.childrens[char-'a'] == nil {return false, word}
        current = current.childrens[char-'a']
        out.WriteRune(char)
        if current.isEnd {
            return true, out.String()
        }
    }
    return false, word
}



func replaceWords(dictionary []string, sentence string) string {
    root := &trieNode{childrens: [26]*trieNode{}}
    for _, word := range dictionary {
        insert(root, word)
    }
    out := new(strings.Builder)
    splitSen := strings.Split(sentence, " ")
    for idx, word := range splitSen {
        if idx != 0 {out.WriteString(" ")}
        found, replace := search(root, word)
        if found {
            splitSen[idx] = replace
        }
        out.WriteString(splitSen[idx])
    }
    return out.String()
}