package dataStruct1

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
		if h.list[idx] < h.list[parentIdx] { //min heap
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

func (h *Heap) Count() int {
	return len(h.list)
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return top
	}

	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1
		if leftIdx >= len(h.list) { //자식 노드가 없다.
			break
		}

		if h.list[leftIdx] < h.list[idx] { //min heap
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) { //오른쪽 자식노드가 있는경우에만
			if h.list[rightIdx] < h.list[idx] { //min heap
				if swapIdx > 0 && h.list[swapIdx] > h.list[rightIdx] { //min heap
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
