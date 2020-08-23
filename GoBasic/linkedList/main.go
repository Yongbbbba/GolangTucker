// 시작하는 node (root라고 부름) 에 대해서 알고 있어야 함
// node를 추가하는 방법은 크게 두 가지가 있는데, 처음부터 쭉 따라가서 tail을 찾는 방법과 tail에 대한 정보를 계속 저장해둬서 그냥 tail 다음에 바로 node를 추가하는 방법이 있음
// 첫 번째 방법은 시간복잡도가 O(N), 후자는 O(1)으로 훨씬 빠름
// node를 삭제하는 방법은 golang의 가비지 컬렉터를 사용할 수 있음.

package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	var tail *Node

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)
	root, tail = RemoveNode(root.next, root, tail) //root 다음 node 지우기

	PrintNodes(root)

	root, tail = RemoveNode(root, root, tail) //root 지우기
	PrintNodes(root)

	root, tail = RemoveNode(tail, root, tail) //tail 지우기
	PrintNodes(root)
	fmt.Println("tail: ", tail.val)
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node //node 추가
	return node
}

func RemoveNode(node, root, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		if root == nil {
			tail = nil
		}

		return root, tail
	}

	prev := root

	for prev.next != node {
		prev = prev.next
	}
	if node == tail {
		prev.next = nil
		tail = prev
	} else {

		prev.next = prev.next.next
	}

	return root, tail
}

// func RemoveNode(prev *Node, root *Node, tail *Node) {
// 	prev.next = prev.next.next //삭제하려고 하는 애를 건너뛴다. 이러면 삭제하려논 node가 가비지 컬렉터에 의해 삭제된다. 참조되고 있지 않기 때문에(reference counter가 0).
// 	// 이 방법은 삭제하고자 하는 node가 root일 경우에는 사용할 수 없기 때문에 root를 없애고 싶으면 그냥 root를 바꾸면 됨
// 	// 반대로 tail을 삭제하기 위해서는 tail을 바꾸면 됨
// }

func PrintNodes(root *Node) {
	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)

}
