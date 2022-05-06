# golang_copy_list_with_random_pointer

A linked list of length `n` is given such that each node contains an additional random pointer, which could point to any node in the list, or `null`.

Construct a **[deep copy](https://en.wikipedia.org/wiki/Object_copying#Deep_copy)** of the list. The deep copy should consist of exactly `n` **brand new** nodes, where each new node has its value set to the value of its corresponding original node. Both the `next` and `random` pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. **None of the pointers in the new list should point to nodes in the original list**.

For example, if there are two nodes `X` and `Y` in the original list, where `X.random --> Y`, then for the corresponding two nodes `x` and `y` in the copied list, `x.random --> y`.

Return *the head of the copied linked list*.

The linked list is represented in the input/output as a list of `n` nodes. Each node is represented as a pair of `[val, random_index]` where:

- `val`: an integer representing `Node.val`
- `random_index`: the index of the node (range from `0` to `n-1`) that the `random` pointer points to, or `null` if it does not point to any node.

Your code will **only** be given the `head` of the original linked list.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2019/12/18/e1.png](https://assets.leetcode.com/uploads/2019/12/18/e1.png)

```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2019/12/18/e2.png](https://assets.leetcode.com/uploads/2019/12/18/e2.png)

```
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

```

**Example 3:**

![https://assets.leetcode.com/uploads/2019/12/18/e3.png](https://assets.leetcode.com/uploads/2019/12/18/e3.png)

```
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]

```

**Constraints:**

- `0 <= n <= 1000`
- $`-10^4$ <= Node.val <=$10^4$`
- `Node.random` is `null` or is pointing to some node in the linked list.

## 解析

這個題目是把一個單向鏈結串列做 deep copy。

這次的鏈結串不同於以往，多了一個 Random pointer 會指向其他在list 內的結點

所以需要避免在複製時，重複複製到已經存在的 Random 指到的結點

要處理這樣的問題可以使用 hashMap 把已存在 Random pointer 紀錄下來

每次發現 hashMap 已經存在就直接拿出來用，不用重複複製

作法是先透過一個 array 依序把目前的結點指標照順序存下來

建立一個 hashMap 來紀錄已存取過的 Random pointer

照 array 逐步檢查 hashMap 是否已經有對應的結點

如果沒有建立一個放到 hashMap 

如果有，存 hashMap 拿出來做指向

這樣實作的 Time Complexity 是 O(n)

Space Complexity 也是 O(n)

## 程式碼

```go
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
```

## 困難點

1. 需要避免重複制已經存在的結點
2. 需要設計一個機制去找詢已經存在的結點

## Solve Point

- [x]  Understand what the problem would like to solve
- [x]  Analysis Complexity