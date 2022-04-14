type TrieNode struct {
    isEnd bool
    childrens [26]*TrieNode
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{
        root: &TrieNode{
            isEnd: false, 
            childrens: [26]*TrieNode{},
        },
    }
}


func (this *Trie) Insert(word string)  {
    cur := this.root
    for i := 0; i < len(word); i++ {
        char := word[i]
        if cur.childrens[char-'a'] == nil {
            cur.childrens[char-'a'] = &TrieNode{
                isEnd: false, childrens: [26]*TrieNode{}}
        }
        cur = cur.childrens[char-'a']
    }
    cur.isEnd = true
}


func (this *Trie) Search(word string) bool {
    cur := this.root
    for i := 0; i < len(word); i++ {
        char := word[i]
        if cur.childrens[char-'a'] == nil {return false}
        cur = cur.childrens[char-'a']
    }
    return cur.isEnd

}


func (this *Trie) StartsWith(prefix string) bool {
    cur := this.root
    for i := 0; i < len(prefix); i++ {
        char := prefix[i]
        if cur.childrens[char-'a'] == nil {return false}
        cur = cur.childrens[char-'a']
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */