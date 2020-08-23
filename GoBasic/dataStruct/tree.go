// root node만 알고 있으면 tree의 다른 node들에 대한 정보 또한 알 수 있다.
// tree는 매우 중요한 자료구조, 자유자재로 쓸 수 있어야 좋다! (뭐든 마찬가지이지만..)
// tree의 순회는 DFS, BFS가 있다.
// DFS는 재귀호출과 스택을 사용
// BFS는 큐를 이용해서 만들 수 있음
// DFS와 BFS의 쓰임새는?
// Tree는 디렉토리 구조에서 사용
// 게임에서 DFS 사용(길찾기), Dijkstra Algorithm
// 노드와 노드사이를 연결하는 것은 모두 그래프라고 한다. 트리와 링크드리스트도 그래프의 한 종류

package dataStruct

import "fmt"

type TreeNode struct {
	Val    int
	Childs []*TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) AddNode(val int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: val}
	} else {
		t.Root.Childs = append(t.Root.Childs, &TreeNode{Val: val})
	}
}

func (t *TreeNode) AddNode(val int) {
	t.Childs = append(t.Childs, &TreeNode{Val: val})
}

//DFS 트리순회
//재귀호출 이용
func DFS1(node *TreeNode) {
	fmt.Printf("%d->", node.Val)

	for i := 0; i < len(node.Childs); i++ {
		DFS1(node.Childs[i])
	}
}

func (t *Tree) DFS1() {
	DFS1(t.Root)

}

//stack을 이용한 방식

func (t *Tree) DFS2() {
	s := []*TreeNode{}
	s = append(s, t.Root)
	for len(s) > 0 {
		var last *TreeNode
		last, s = s[len(s)-1], s[:len(s)-1]

		fmt.Printf("%d->", last.Val)
		for i := 0; i < len(last.Childs); i++ {
			s = append(s, last.Childs[i])
		}
	}
}

// BFS 너비우선탐색
func (t *Tree) BFS() {
	queue := []*TreeNode{}
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		var first *TreeNode
		first, queue = queue[0], queue[1:]

		fmt.Printf("%d->", first.Val)
		for i := 0; i < len(first.Childs); i++ {
			queue = append(queue, first.Childs[i])
		}
	}

}
