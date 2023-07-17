package iterables

type Iterable[T any] struct {
	gen Generator[T]
}

type IterationStop struct {
}

type Generator[T any] interface {
	Next() (T, error)
}

type invalidIterable[T any] struct {
	err error
}

type generator[T any, S any] struct {
	state       S
	genFunction func(*S) (T, error)
}

func (g *generator[T, S]) Next() (T, error) {

	return g.genFunction(&g.state)
}

func (i *invalidIterable[T]) Next() (T, error) {

	return *new(T), i.err
}

func (i Iterable[T]) Next() (T, error) {

	return i.gen.Next()
}

func (i IterationStop) Error() string {
	return "The iteration stopped"
}
