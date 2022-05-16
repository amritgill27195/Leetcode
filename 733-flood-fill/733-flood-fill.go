func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    oldColor := image[sr][sc]
    if oldColor == newColor {return image}
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    m := len(image)
    n := len(image[0])
    
    q := new(queue)
    image[sr][sc] = newColor
    q.enqueue(sr,sc)

    for !q.isEmpty() {
        r,c := q.dequeue()
        for _, dir := range dirs {
            nr := r+dir[0]
            nc := c+dir[1]
            if nr >= 0 && nr < m && nc >=0 && nc < n && image[nr][nc] == oldColor {
                image[nr][nc] = newColor
                q.enqueue(nr,nc)
            }
        }
    }
    
    return image
}


type listNode struct {
    r,c int
    next *listNode
}
type queue struct {
    head *listNode
    tail *listNode
}


func (q *queue) isEmpty() bool {
    return q.head == nil
}
func (q *queue) enqueue(r,c int) {
    newNode := &listNode{r:r,c:c}
    if q.head == nil {
        q.head = newNode
        q.tail = newNode
        return
    }
    q.tail.next = newNode
    q.tail = newNode
}

func (q *queue) dequeue() (int,int) {
    if q.head == nil {
        return -1,-1
    }
    r := q.head.r
    c := q.head.c
    
    if q.head.next == nil {
        q.head = nil
        q.tail = nil
        return r,c
    }
    
    newHead := q.head.next
    q.head.next = nil 
    q.head = newHead
    return r,c
}