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
    q := []*TreeNode{root}
    strBuilder := new(strings.Builder)
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        if dq == nil {
            strBuilder.WriteString("null")
        } else {
            strBuilder.WriteString(fmt.Sprintf("%v",dq.Val))
            q = append(q, dq.Left)
            q = append(q, dq.Right)
        }
        strBuilder.WriteString(",")
    }
    return strBuilder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data == "" {return nil}
    strData := strings.Split(data, ",")
    idxZeroInt, _ := strconv.Atoi(strData[0])
    root := &TreeNode{Val:idxZeroInt}
    idx := 1
    q := []*TreeNode{root}
    
    for len(q) != 0 && idx < len(strData)  {
        dq := q[0]
        q = q[1:]
        if strData[idx] != "null" {
            intVal, _ := strconv.Atoi(strData[idx])
            dq.Left = &TreeNode{Val: intVal}
            q = append(q, dq.Left)
        }
        idx++
        if idx < len(strData) {
            if strData[idx] != "null" {
                intVal, _ := strconv.Atoi(strData[idx])
                dq.Right = &TreeNode{Val: intVal}
                q = append(q, dq.Right)
            }
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