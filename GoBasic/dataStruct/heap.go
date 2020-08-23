// 최대 힙과 최소 힙으로 나뉘어짐
// 최대 힙이란, 부모노드는 항상 자식노드보다 값이 큰 힙을 말함
// 힙이 쓰이는 곳 : 최대값과 최소값을 찾아야할 때
// priority queue를 heap으로 만들 수 있음
// 힙을 이용해서 정렬하는 것을 힙 정렬이라고 함
// 힙의 속도는 층수와 비례한다
// 힙을 이용하면 최소신장트리가 자동으로 만들어진다.
// 보통 힙은 트리가 아니고 배열을 통해서 구현을 많이 한다. 배열을 사용하면 끝에 노드가 무엇인지 알기 쉽다.

package dataStruct

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)
	idx := len(h.list) - 1

	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}

}

func (h *Heap) Print() {
	fmt.Println(h.list)
}

func (h *Heap) Pop() int {

	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1
		if leftIdx >= len(h.list) {
			break // 자식노드가 없다는 뜻
		}
		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}
		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}
		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}

// heap을 이용한 알고리즘 문제풀이 문제 예시
// 안녕하세요, 매일프로그래밍 이번 주 문제입니다.
// 정수 배열(int array)과 정수 N이 주어지면, N번째로 큰 배열 원소를 찾으시오.
// 예제) input : [-1,3,-1,5,4], 2 , output: 4
