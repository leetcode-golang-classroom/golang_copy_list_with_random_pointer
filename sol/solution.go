package sol

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	arr := []*Node{}
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	var copyHead *Node
	var curN *Node
	randomMap := make(map[*Node]*Node)
	for idx, node := range arr {
		if idx == 0 {
			copyHead = &Node{Val: node.Val}
			randomMap[node] = copyHead
			if node.Random != nil {
				if random, exists := randomMap[node.Random]; exists {
					copyHead.Random = random
				} else {
					randomMap[node.Random] = &Node{Val: node.Random.Val}
					copyHead.Random = randomMap[node.Random]
				}
			}
			curN = copyHead
		} else {
			if n, exists := randomMap[node]; exists {
				curN.Next = n
			} else {
				randomMap[node] = &Node{Val: node.Val}
				curN.Next = randomMap[node]
			}
			if node.Random != nil {
				if random, ok := randomMap[node.Random]; ok {
					curN.Next.Random = random
				} else {
					randomMap[node.Random] = &Node{Val: node.Random.Val}
					curN.Next.Random = randomMap[node.Random]
				}
			}
			curN = curN.Next
		}
	}

	return copyHead
}
