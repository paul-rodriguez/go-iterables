package iterables

func (head Iterable[T]) Extend(tail Iterable[T]) Iterable[T] {

	gen := extend[T]{head, tail, false}
	return Iterable[T]{&gen}
}

type extend[T any] struct {
	source0   Iterable[T]
	source1   Iterable[T]
	exhausted bool
}

func (e *extend[T]) Next() (T, error) {

	if !e.exhausted {
		item, err := e.source0.Next()
		if (err == IterationStop{}) {
			e.exhausted = true
		} else if err != nil {
			return *new(T), err
		} else {
			return item, nil
		}
	}

	item, err := e.source1.Next()
	if err != nil {
		return *new(T), err
	}
	return item, nil
}
