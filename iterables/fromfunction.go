package iterables

func FromFunction[T any](origin func() (T, error)) Iterable[T] {

	gen := fromFunction[T]{origin}
	return Iterable[T]{&gen}
}

type fromFunction[T any] struct {
	origin func() (T, error)
}

func (gen *fromFunction[T]) Next() (T, error) {

	return gen.origin()
}
