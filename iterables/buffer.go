package iterables

import "sync"

type Buffer[T any] struct {
	front *list[T]
	back  *list[T]
	mutex sync.Mutex
}

func (b *Buffer[T]) Next() (T, error) {

	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.IsEmpty() {
		return *new(T), EmptyBufferError{}
	} else {
		return b.pop()
	}
}

func (b *Buffer[T]) Load(t T, err error) {

	b.mutex.Lock()
	defer b.mutex.Unlock()

	next := b.front
	newList := list[T]{next: next, prev: nil, item: t, err: err}
	if next != nil {
		next.prev = &newList
	} else {
		b.back = &newList
	}
	b.front = &newList
}

func (b *Buffer[T]) IsEmpty() bool {

	return b.front == nil
}

type EmptyBufferError struct {
}

func (e EmptyBufferError) Error() string {

	return "the buffer was empty"
}

type list[T any] struct {
	next *list[T]
	prev *list[T]
	item T
	err  error
}

func (b *Buffer[T]) pop() (T, error) {
	result := b.back.item
	err := b.back.err
	previous := b.back.prev
	b.back = previous
	if previous != nil {
		previous.next = nil
	} else {
		b.front = nil
	}
	return result, err
}
