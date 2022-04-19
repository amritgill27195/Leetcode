
type trieNode struct {
    isEnd bool
    childrens [26]*trieNode
}

func insert(word string, root *trieNode) {
    curr := root
    for i := 0; i < len(word); i++ {
        if curr.childrens[word[i]-'a'] == nil {
            curr.childrens[word[i]-'a'] = new(trieNode)
        }
        curr = curr.childrens[word[i]-'a']
    }
    curr.isEnd = true
}

func search(word string, root *trieNode) (bool, string) {
    curr := root
    out := new(strings.Builder)
    for i := 0; i < len(word); i++ {
        if curr.childrens[word[i]-'a'] == nil { return false, word}
        out.WriteByte(word[i])
        curr = curr.childrens[word[i]-'a']
        if curr.isEnd {return true, out.String()}
    }
    return false, word
}


func replaceWords(dictionary []string, sentence string) string {
    
    root := new(trieNode)
    for _, word := range dictionary{
        insert(word, root)
    }
    out := new(strings.Builder)
    splitSent := strings.Split(sentence, " ")
    for idx, word := range splitSent{
        if idx != 0 {out.WriteString(" ")}
        found, replace := search(word, root)
        if found {
           splitSent[idx] = replace 
        }
        out.WriteString(splitSent[idx])
    }
    return out.String()
}