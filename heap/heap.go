package heap
import (
	"errors"
)

// Heap  Common heap interface.
// May have more diffrent types of heap implements.
type Heap interface {
	// Basic heap operations
	Initialize(capacity int, f compareFunc) Heap
	Add(ele interface{}) error
	GetMin() (interface{}, error)
	DeleteMin() (interface{}, error)
}

var (
	errHeapEmpty = errors.New("heap is empty")
)

type compareFunc func (a, b interface{}) bool

type heap struct {
	size int
	capacity int
	elements []interface{}
	compare compareFunc
}

type min struct{}
func(m *min) IsMin() bool {
	return true
}

func NewHeap(capacity int, f compareFunc) Heap {
	h := &heap{}
	return h.Initialize(capacity, f)
}

func BuildHeap(capacity int, eles []interface{}, f compareFunc) Heap {
	h := &heap{}
	h.size = len(eles)
	h.capacity = capacity
	h.elements= make([]interface{}, capacity + 1)
	h.elements[0] = min{}
	for i, e := range eles {
		h.elements[i+1] = e
	}
	h.compare = f
	for i:=h.size/2;i>0;i-- {
		h.percolateDown(i)
	}
	return h
}

func (h *heap) cmp(a, b interface{}) bool {
	if _, ok := a.(min);ok {
		return false
	}
	if _, ok := b.(min);ok {
		return true
	}
	return h.compare(a, b)
}

func (h *heap) Initialize(capacity int, f compareFunc) Heap {
	h.capacity = capacity
	h.size= 0
	h.elements= make([]interface{}, capacity + 1)
	h.compare=f
	h.elements[0] = min{}
	return h
}

func (h *heap) parent(index int) interface{} {
	return h.elements[index/2]
}

func (h *heap) Add(ele interface{}) error {
	if h.size >= h.capacity {
		return errors.New("Heap is full")
	}
	var index int
	for index = h.size + 1; h.cmp(h.parent(index),ele);index = index/2{
		h.elements[index] = h.elements[index/2]
	}
	h.elements[index]= ele
	h.size = h.size + 1
	return nil
}

func (h *heap) GetMin() (interface{},error) {
	if h.size == 0 {
		return nil, errHeapEmpty
	}
	return h.elements[1], nil
}

func (h *heap) DeleteMin() (interface{}, error) {
	if h.size == 0 {
		return nil, errHeapEmpty
	}
	min := h.elements[1]
	x := h.elements[h.size]
	var index, child int
	for index = 1;index*2 <=h.size; {
		if index*2+1 <= h.size && h.cmp(h.elements[index*2], h.elements[index*2 + 1]) {
			child = index * 2 + 1
		} else {
			child = index * 2
		}
		if h.cmp(h.elements[child],x) {
			break
		}
		h.elements[index] = h.elements[child]
		index = child
	}
	h.elements[index] = x
	h.size = h.size - 1
	return min, nil
}

func (h *heap) percolateDown(i int) {
	if h.size == 0 {
		return 
	}
	var index, child int
	x := h.elements[i]
	for index = i;index*2 <=h.size; {
		if index*2+1 <= h.size && h.cmp(h.elements[index*2], h.elements[index*2 + 1]) {
			child = index * 2 + 1
		} else {
			child = index * 2
		}
		if h.cmp(h.elements[child],x) {
			break
		}
		h.elements[index] = h.elements[child]
		index = child
	}
	h.elements[index] = x
}