package iterables

func (i Iterable[T]) Take(n int) Iterable[T] {

	gen := take[T]{n, i}
	return Iterable[T]{&gen}
}

type take[T any] struct {
	remaining int
	source    Iterable[T]
}

func (gen *take[T]) Next() (T, error) {

	if gen.remaining <= 0 {
		return *new(T), IterationStop{}
	}

	gen.remaining--
	return gen.source.Next()
}
