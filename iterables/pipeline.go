package iterables

func (i Iterable[T]) Pipeline() Iterable[T] {

	ch := make(chan result[T])

	go func() {
		for {
			item, err := i.Next()
			if err != nil {
				ch <- result[T]{item, err}
				close(ch)
				return
			}
			ch <- result[T]{item, err}
		}
	}()

	gen := pipeline[T]{ch, false}
	return Iterable[T]{&gen}
}

type pipeline[T any] struct {
	ch      chan result[T]
	stopped bool
}

func (gen *pipeline[T]) Next() (T, error) {

	if gen.stopped {
		return *new(T), IterationStop{}
	}
	result := <-gen.ch
	if result.err != nil {
		gen.stopped = true
	}
	return result.item, result.err
}
