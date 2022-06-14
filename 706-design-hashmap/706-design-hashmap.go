const (
    arraySize = 1001
)

type MyHashMap struct {
    items [][]*int   
}


func Constructor() MyHashMap {
    return MyHashMap{
        items: make([][]*int, arraySize),
    }  
}

func (this *MyHashMap) hashIdx(key int) int {
    return key % arraySize
}

func (this *MyHashMap) hashIdx2(key int) int {
    return key / arraySize
}


func (this *MyHashMap) Put(key int, value int)  {
    idx := this.hashIdx(key)
    if this.items[idx] == nil {
        this.items[idx] = make([]*int, arraySize)
    }
    innerIdx := this.hashIdx2(key)
    this.items[idx][innerIdx] = &value
}


func (this *MyHashMap) Get(key int) int {
    idx := this.hashIdx(key)
    innerIdx := this.hashIdx2(key)
    if this.items[idx] == nil || this.items[idx][innerIdx] == nil {
        return -1
    }
    return *this.items[idx][innerIdx]
}


func (this *MyHashMap) Remove(key int)  {
    idx := this.hashIdx(key)
    innerIdx := this.hashIdx2(key)
    if this.items[idx] == nil {return}
    this.items[idx][innerIdx] = nil
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */