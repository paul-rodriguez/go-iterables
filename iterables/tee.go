package iterables

import "sync"

func (i Iterable[T]) Tee() (Iterable[T], Iterable[T]) {

	buffers := []Buffer[T]{{}, {}}
	splitter := tee[T]{source: i, buffers: buffers}

	connector0 := teeConnector[T]{&splitter, 0}
	connector1 := teeConnector[T]{&splitter, 1}
	iter0 := Iterable[T]{&connector0}
	iter1 := Iterable[T]{&connector1}
	return iter0, iter1
}

type tee[T any] struct {
	source  Iterable[T]
	buffers []Buffer[T]
	mutex   sync.Mutex
}

func (t *tee[T]) takeFrom(id int) (T, error) {

	t.mutex.Lock()
	defer t.mutex.Unlock()

	buf := &t.buffers[id]
	if buf.IsEmpty() {
		item, err := t.source.Next()
		for i := range t.buffers {
			t.buffers[i].Load(item, err)
		}
	}
	item, err := buf.Next()
	return item, err
}

type teeConnector[T any] struct {
	source *tee[T]
	id     int
}

func (t *teeConnector[T]) Next() (T, error) {

	return t.source.takeFrom(t.id)
}
