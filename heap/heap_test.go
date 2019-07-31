package heap

import (
	"testing"
	"fmt"
)

func buildHeapByAdd() Heap {
	h := NewHeap(10, func(a,b interface{}) bool {
		return a.(int) > b.(int)
	})
	/*
	h.Add(10)
	h.Add(2)
	h.Add(4)
	h.Add(8)
	h.Add(1)
	*/
	return h
}

//TODO: use assert
func TestHeapAdd(t *testing.T) {
	h := NewHeap(10, func(a,b interface{}) bool {
		return a.(int) > b.(int)
	})

	h.Add(10)
	h.Add(2)
	h.Add(4)
	h.Add(8)
	h.Add(1)
	m, err := h.GetMin()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
}

func TestDeleteMin(t *testing.T) {
	h := buildHeapByAdd()
	m, err := h.DeleteMin()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
	m, err = h.GetMin()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}

func TestBuildHeap(t *testing.T) {
	eles := []interface{}{10,6,4,5,8,9,1}
	h := BuildHeap(20, eles,func(a,b interface{}) bool {
		return a.(int) > b.(int)
	})
	m, err := h.GetMin()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(m)
}