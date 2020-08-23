// double likedlist의 concept를 잘 기억해서 다시 구현해보는 것이 중요
// 이것과 똑같이 만들 필요는 없고 concept만 내 방법으로 구현하는 것이 중요하다
// 특히 이번 강의는 날림으로 들었기 때문에 내가 나중에 만들어보는 것이 좋다.. 

// 슬라이스(또는 배열)과 리스트의 차이
// 슬라이스에 새로운 요소를 추가할 때 capacity가 없으면 O(N) (capacity가 있을 때는 O(1))
// 리스트는 O(1), 그냥 tail 뒤에다가 붙이면 되니까
// 슬라이스의 요소를 삭제할 때, 맨 뒤에 요소를 삭제하면 O(1)이지만 중간에 있는 것을 삭제할 때는 복사를 이용해야함. slicing, append를 이용, O(N)
// 그럼에도 슬라이스를 쓰는 이유? 
// 슬라이스는 시작 메모리 주소를 알고, 사이즈를 알고 있기 때문에 가져오고 싶어하는 특정인덱스의 메모리 주소를 즉각 알 수 있다. 
// 반면 리스트는 중간에 있는 것을 가져오기 위해 root부터 해당 요소까지 계속 전진해야함
// 즉, Random Access에서 slice가 유리 
// 또한, 리스트에 비해 cache miss를 최소화할 수 있다.





package main

import "fmt"

type Node struct {
 next *Node
 prev *Node // 이전 노드에 대해서도 기억
 val  int
}

type LinkedList struct {
 root *Node
 tail *Node
}

func (l *LinkedList) AddNode(val int) {
 if l.root == nil {
  l.root = &Node{val: val}
  l.tail = l.root
  return
 }
 l.tail.next = &Node{val: val}
 prev := l.tail
 l.tail = l.tail.next
 l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
 if node == l.root {
  l.root = l.root.next
  l.root.prev = nil
  node.next = nil
  return
 }

 prev := node.prev

 if node == l.tail {
  prev.next = nil
  l.tail.prev = nil
  l.tail = prev
 } else {
  node.prev = nil
  prev.next = prev.next.next
  prev.next.prev = prev
 }
 node.next = nil
}

func (l *LinkedList) PrintNodes() {
 node := l.root
 for node.next != nil {
  fmt.Printf("%d -> ", node.val)
  node = node.next
 }
 fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverse() {
 node := l.tail
 for node.prev != nil {
  fmt.Printf("%d -> ", node.val)
  node = node.prev
 }
 fmt.Printf("%d\n", node.val)
}

func main() {
 list := &LinkedList{}
 list.AddNode(0)

 for i := 1; i < 10; i++ {
  list.AddNode(i)
 }

 list.PrintNodes()

 list.RemoveNode(list.root.next)

 list.PrintNodes()

 list.RemoveNode(list.root)

 list.PrintNodes()

 list.RemoveNode(list.tail)

 list.PrintNodes()
 fmt.Printf("tail:%d\n", list.tail.val)

 list.PrintReverse()

 a := []int{1, 2, 3, 4, 5}
 a = append(a[0:2], a[3:]...)
 fmt.Println(a)
 
}