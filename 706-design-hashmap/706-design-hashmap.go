const (
    arraySize = 1000000
)

type MyHashMap struct {
    items []*int   
}


func Constructor() MyHashMap {
    return MyHashMap{
        items: make([]*int, arraySize),
    }  
}

func (this *MyHashMap) hashIdx(key int) int {
    return key % arraySize
}


func (this *MyHashMap) Put(key int, value int)  {
    idx := this.hashIdx(key)
    this.items[idx] = &value
}


func (this *MyHashMap) Get(key int) int {
    idx := this.hashIdx(key)
    if this.items[idx] == nil {
        return -1
    }
    return *this.items[idx]
}


func (this *MyHashMap) Remove(key int)  {
    idx := this.hashIdx(key)
    this.items[idx] = nil
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */