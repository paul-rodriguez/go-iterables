package iterables

func Apply[T any, U any](i Iterable[T], f func(T) (U, error)) Iterable[U] {

	gen := apply[T, U]{i, f}
	return Iterable[U]{&gen}
}

type apply[T any, U any] struct {
	source    Iterable[T]
	transform func(T) (U, error)
}

func (a *apply[T, U]) Next() (U, error) {

	item, err := a.source.Next()
	if err != nil {
		return *new(U), err
	}
	result, err := a.transform(item)
	if err != nil {
		return *new(U), err
	}
	return result, nil
}
