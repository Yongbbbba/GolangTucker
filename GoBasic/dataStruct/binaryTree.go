// 이진트리
// 이진트리는 자식 노드가 두 개밖에 없는 경우
// 왼쪽은 그 노드보다 value가 작은 노드, 오른쪽은 그 노드보다 value가 큰 노드가 위치
// 이진트리가 중요한 이유는 BST(Binary Search Tree) 때문
// 검색을 빨리 하기 위해 사용
// 최소 신장 트리가 가장 작은 키를 갖기 때문에 가장 효율적인 트리이다.
// 최소신장트리를 만들기 위해서 몇 가지 방법을 사용하는데, 그중에 AVL 트리가 가장 대표적인 트리

package dataStruct

import "fmt"

type BinaryTreeNode struct {
	Val int

	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if n.Val > v {
		if n.left == nil {
			n.left = &BinaryTreeNode{Val: v}
			return n.left
		} else {
			return n.left.AddNode(v)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{Val: v}
			return n.right
		} else {
			return n.right.AddNode(v)
		}
	}

}

type depthNode struct {
	depth int
	node  *BinaryTreeNode
}

func (t *BinaryTree) Print() {
	// BFS를 이용하여 순회
	// 큐 사용
	q := []depthNode{}
	q = append(q, depthNode{depth: 0, node: t.Root})

	currentDepth := 0

	for len(q) > 0 {
		var first depthNode
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Print(first.node.Val, " ")

		if first.node.left != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.left})
		}
		if first.node.right != nil {
			q = append(q, depthNode{depth: currentDepth + 1, node: first.node.right})

		}
	}

}

func (t *BinaryTree) Search(v int) (bool, int) {
	return t.Root.Search(v, 1)

}

func (n *BinaryTreeNode) Search(v int, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.left != nil {
			return n.left.Search(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.right != nil {
			return n.right.Search(v, cnt+1)
		}
		return false, cnt
	}

}
