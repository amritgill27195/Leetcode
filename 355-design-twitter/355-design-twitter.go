
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

// time: o(nm) 
// space: o(n)
func (this *Twitter) GetNewsFeed(userId int) []int {

    usersSet := this.followMap[userId]    
    if usersSet == nil {return nil}
    
    
    // time: o(n) * o(m) * o(10log10) -- where n is the number of users and m is the number of tweets for each nth user
    mh := &minHeap{tweets: []*tweet{}}
    // at worse we have n users
    for uidKey, _ := range usersSet.items {
        // For each user, we only care about the latest 10 so start from back and count 10
        tweets := this.tweetMap[uidKey]
        tSize := len(tweets)
        for i := tSize-1; i >= 0 && i >= tSize-11; i-- {
            heap.Push(mh, tweets[i])
            if len(mh.tweets) > 10 {
                heap.Pop(mh)
            }
        }
    }
    
    // time: worse case we have 10 tweets to output - so this loop at worse will be o(10) or in other words constant in worst case
    // space : this is generating output, so we do not count output array as space
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