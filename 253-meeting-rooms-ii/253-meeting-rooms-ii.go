
type MinHeap struct {
	items []int
}

func (mn *MinHeap) insert(val int) {
	mn.items = append(mn.items, val)
	mn.heapifyUp(len(mn.items) - 1)
}

func (mn *MinHeap) extract() int {
	l := len(mn.items)-1
	out := mn.items[0]
	mn.swap(0, l)
	mn.items = mn.items[:l]
	mn.heapifyDown(0)
	return out
}

func (mn *MinHeap) swap(a,b int) {
	mn.items[a], mn.items[b] = mn.items[b], mn.items[a]
}

func (mn *MinHeap) heapifyDown(idx int) {
	l := len(mn.items)-1
	if idx < 0 || idx > l {
		return
	}
	left := (2*idx) +1
	right := (2*idx) +2
	for (left <= l && mn.items[left] < mn.items[idx]) ||
		(right <= l && mn.items[right] < mn.items[idx]) {
		swappingWith := idx
		if left <= l && mn.items[left] < mn.items[swappingWith] {
			swappingWith = left
		}
		if right <= l && mn.items[right] < mn.items[swappingWith] {
			swappingWith = right
		}
		if swappingWith == idx {
			return
		}
		mn.swap(swappingWith, idx)
		idx = swappingWith
		left = (2*idx) +1
		right = (2*idx) +2
	}
}

func (mn *MinHeap) heapifyUp(idx int) {
	l := len(mn.items)-1
	if idx < 0 || idx > l {
		return
	}
	parent := int(math.Floor(float64((idx-1)/2)))
	for parent <= l && mn.items[idx] < mn.items[parent]{
		mn.swap(parent, idx)
		idx = parent
		parent = int(math.Floor(float64((idx-1)/2)))
	}
}


func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
// 		// edge cases
// 		if len(intervals[i]) == 0 && len(intervals[j]) == 0 {
// 			return false // two empty slices - so one is not less than other i.e. false
// 		}
// 		if len(intervals[i]) == 0 || len(intervals[j]) == 0 {
// 			return len(intervals[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
// 		}

// 		// both slices len() > 0, so can test this now:
		return intervals[i][0] < intervals[j][0]
	})
	mn := &MinHeap{items: []int{}}

	for _, e := range intervals {
		startTime := e[0]
		endTime := e[1]
		if len(mn.items) == 0  {
			mn.insert(endTime)
		} else {
			earliestEndTime := mn.items[0]
			if earliestEndTime <= startTime {
				_ = mn.extract()
			}
			mn.insert(endTime)
		}
	}
	return len(mn.items)
}

