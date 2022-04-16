
type Twitter struct {
    followMap map[int]*set
    tweetMap map[int][]*tweet
}

type tweet struct {
    id int
    t time.Time
}

func Constructor() Twitter {
    return Twitter{
        followMap: map[int]*set{},
        tweetMap: map[int][]*tweet{},
    }
}

// time: o(1) 
// space: o(1)
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    if this.followMap[userId] == nil {
        this.followMap[userId] = newSet()
    }
    this.followMap[userId].add(userId)    
    this.tweetMap[userId] = append(this.tweetMap[userId], &tweet{id: tweetId, t: time.Now()})
}

// time: o(n) 
// space: o(n) -
func (this *Twitter) GetNewsFeed(userId int) []int {

    // at worse we have this user following n users
    // space: o(n)
    // time: o(1)
    users := this.followMap[userId].list()
    
    
    // min heap at worse will be storing 10 tweet objs, which is constant
    mh := &minHeap{tweets: []*tweet{}}
    // at worse we have n users
    for _, uid := range users {
        // at worse we have m tweets for each n user
        // there time: o(mn)
        for _, tid := range this.tweetMap[uid] {
            heap.Push(mh, tid)
            if len(mh.tweets) > 10 {
                heap.Pop(mh)
            }
        }
    }
    
    out := make([]int, len(mh.tweets))
    for i := len(mh.tweets)-1; i>=0; i-- {
        out[i] = heap.Pop(mh).(*tweet).id
    }
    return out
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.followMap[followerId] == nil {
        this.followMap[followerId] = newSet()
    }
    this.followMap[followerId].add(followeeId)
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.followMap[followerId] == nil {
        return
    }
    this.followMap[followerId].remove(followeeId)
}



// Heap -- implements container heap interface : https://pkg.go.dev/container/heap#Interface
type minHeap struct {
	tweets []*tweet
}
func (m *minHeap) Len() int {return len(m.tweets)}
func (m *minHeap) Less(i, j int) bool {return m.tweets[i].t.Before(m.tweets[j].t)}
func (m *minHeap) Swap(i, j int) { m.tweets[i],m.tweets[j] = m.tweets[j], m.tweets[i]}
func (m *minHeap) Push(x interface{}) {m.tweets = append(m.tweets, x.(*tweet))}
func (m *minHeap) Pop() interface{} {
	i := m.tweets[len(m.tweets)-1]
	m.tweets = m.tweets[:len(m.tweets)-1]
	return i
}


// Poor mans set 
type set struct {
    items map[int]struct{}
}
func newSet() *set {
    return &set{items:map[int]struct{}{}}
}
func (s *set) add(item int) {
    s.items[item] = struct{}{}
}
func (s *set) remove(item int){
    delete(s.items, item)
}
func (s *set) list() []int{
    if s == nil || s.items == nil {return nil}
    out := []int{}
    for k, _ := range s.items{
        out = append(out, k)
    }
    return out
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */