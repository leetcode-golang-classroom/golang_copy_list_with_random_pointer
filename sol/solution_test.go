package sol

import (
	"reflect"
	"testing"
)

func CreateList(list *[][]int) *Node {
	arr := *list
	var head *Node
	var cur *Node
	pointerMap := make(map[int]*Node)
	for idx, node := range arr {
		if idx == 0 {
			head = &Node{Val: node[0]}
			pointerMap[idx] = head
			cur = head
		} else {
			pointerMap[idx] = &Node{Val: node[0]}
			cur.Next = pointerMap[idx]
			cur = cur.Next
		}
	}
	cur = head
	for _, node := range arr {
		if len(node) > 1 {
			if val, exists := pointerMap[node[1]]; exists {
				cur.Random = val
			}
		}
		cur = cur.Next
	}
	return head
}
func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "head = [[7,null],[13,0],[11,4],[10,2],[1,0]]",
			args: args{head: CreateList(&[][]int{{7}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})},
			want: CreateList(&[][]int{{7}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}),
		},
		{
			name: "head = [[1,1],[2,1]]",
			args: args{head: CreateList(&[][]int{{1, 1}, {2, 1}})},
			want: CreateList(&[][]int{{1, 1}, {2, 1}}),
		},
		{
			name: "head = [[3,null],[3,0],[3,null]]",
			args: args{head: CreateList(&[][]int{{3}, {3, 0}, {3}})},
			want: CreateList(&[][]int{{3}, {3, 0}, {3}}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
