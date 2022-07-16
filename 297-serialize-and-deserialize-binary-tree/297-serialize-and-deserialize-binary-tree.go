/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    out := new(strings.Builder)
    q := []*TreeNode{root}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == nil {
            out.WriteString("null")
        } else {
            out.WriteString(fmt.Sprintf("%v", dq.Val))
        }
        if dq != nil {
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        }
        if len(q) != 0 {
            out.WriteString(",")
        }
    }
    return out.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data == "" {return nil}
    sliceData := strings.Split(data, ",")
    rootInt, _ := strconv.Atoi(sliceData[0])
    root := &TreeNode{Val: rootInt}
    i := 1
    q := []*TreeNode{root}
    for i < len(sliceData) && len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if sliceData[i] != "null" {
            leftInt, _ := strconv.Atoi(sliceData[i])
            left := &TreeNode{Val: leftInt}
            dq.Left = left
            q = append(q, left)
        }
        i++
        if sliceData[i] != "null" {
            rightInt, _ := strconv.Atoi(sliceData[i])
            right := &TreeNode{Val: rightInt}
            dq.Right = right
            q = append(q, right)            
        }
        i++
    }
    return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */