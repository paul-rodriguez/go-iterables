package iterables

func (i Iterable[T]) Filter(f func(T) bool) Iterable[T] {

	g := filter[T]{i, f}
	return Iterable[T]{&g}
}

type filter[T any] struct {
	source    Iterable[T]
	predicate func(T) bool
}

func (g *filter[T]) Next() (T, error) {

	for true {
		item, err := g.source.Next()
		if err != nil {
			return *new(T), err
		}
		if g.predicate(item) {
			return item, nil
		}
	}
	return *new(T), IterationStop{}
}
