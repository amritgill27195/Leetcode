const (
    arraySize = 1000
    innerArraySize = arraySize+1
)

type MyHashMap struct {
    items [][]*int
}

func hashIdx(key int) int {
    return key % arraySize
}
func toPtrInt(n int) *int {return &n}

func doubleHash(key int) int {
    return key / innerArraySize
}


func Constructor() MyHashMap {
    return MyHashMap{
        items: make([][]*int, arraySize),
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    hashIdx := hashIdx(key)
    hashIdx2 := doubleHash(key)
    if this.items[hashIdx] == nil {
        this.items[hashIdx] = make([]*int, innerArraySize)
    }
    this.items[hashIdx][hashIdx2] = toPtrInt(value)
}


func (this *MyHashMap) Get(key int) int {
    hashIdx := hashIdx(key)
    hashIdx2 := doubleHash(key)
    if this.items[hashIdx] == nil || this.items[hashIdx][hashIdx2] == nil {
        return -1
    }
    return *this.items[hashIdx][hashIdx2]
}


func (this *MyHashMap) Remove(key int)  {
    hashIdx := hashIdx(key)
    hashIdx2 := doubleHash(key)
    if this.items[hashIdx] == nil {
        return
    }
    this.items[hashIdx][hashIdx2] = nil
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */