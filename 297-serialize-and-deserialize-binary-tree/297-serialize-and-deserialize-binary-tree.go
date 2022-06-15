/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    str := new(strings.Builder)
    q := []*TreeNode{root}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == nil {
            str.WriteString("null")
        } else {
            str.WriteString(fmt.Sprintf("%v", dq.Val))
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        }
        str.WriteString(",")
    }
    
    return str.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data == "" {return nil}
    strData := strings.Split(data, ",")
    intVal,_ := strconv.Atoi(strData[0])
    root := &TreeNode{Val: intVal}
    idx := 1
    q := []*TreeNode{root}
    
    for len(q) != 0 && idx < len(strData) {
        dq := q[0]
        q = q[1:]
        if strData[idx] != "null" {
            intVal,_ := strconv.Atoi(strData[idx])
            dq.Left = &TreeNode{Val: intVal}
            q = append(q, dq.Left)
        }
        idx++
        if idx < len(strData) && strData[idx] != "null" {
            intVal,_ := strconv.Atoi(strData[idx])
            dq.Right = &TreeNode{Val: intVal}
            q = append(q, dq.Right)
        }
        idx++
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