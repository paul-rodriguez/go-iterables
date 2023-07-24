package iterables

import (
	"container/heap"
	"fmt"
)

func (i Iterable[T]) Sort(cmp func(T, T) int) Iterable[T] {

	gen := sort[T]{tHeap[T]{cmpFunc: cmp}}
	for true {
		item, err := i.Next()
		if (err == IterationStop{}) {
			break
		} else if err != nil {
			invalid := invalidIterable[T]{InvalidSort{err}}
			return Iterable[T]{&invalid}
		}
		heap.Push(&gen.sorter, item)
	}
	return Iterable[T]{&gen}
}

type sort[T any] struct {
	sorter tHeap[T]
}

func (s *sort[T]) Next() (T, error) {

	if s.sorter.Len() <= 0 {
		return *new(T), IterationStop{}
	}
	popped := heap.Pop(&s.sorter).(T)
	return popped, nil
}

type tHeap[T any] struct {
	cmpFunc func(T, T) int
	data    []T
}

func (h *tHeap[T]) Push(x any) {

	item := x.(T)
	h.data = append(h.data, item)
}

func (h *tHeap[T]) Pop() any {

	lastIndex := len(h.data) - 1
	result := h.data[lastIndex]
	h.data = h.data[:lastIndex]
	return result
}

func (h *tHeap[T]) Len() int {

	return len(h.data)
}

func (h *tHeap[T]) Less(i int, j int) bool {

	itemI := h.data[i]
	itemJ := h.data[j]

	return h.cmpFunc(itemI, itemJ) < 0
}

func (h *tHeap[T]) Swap(i int, j int) {

	save := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = save
}

type InvalidSort struct {
	Cause error
}

func (i InvalidSort) Error() string {
	return fmt.Sprintf(
		"invalid sort because underlying iterable failed: %v",
		i.Cause)
}
