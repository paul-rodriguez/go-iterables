package iterables

type Pair[T any, U any] struct {
	F0 T
	F1 U
}

func (p Pair[T, U]) Items() (T, U) {
	return p.F0, p.F1
}

func Zip2[T any, U any](ts Iterable[T], us Iterable[U]) Iterable[Pair[T, U]] {

	gen := zip2[T, U]{ts, us}
	return Iterable[Pair[T, U]]{&gen}
}

type zip2[T any, U any] struct {
	source0 Iterable[T]
	source1 Iterable[U]
}

func (z *zip2[T, U]) Next() (Pair[T, U], error) {

	item0, err0 := z.source0.Next()
	if err0 != nil {
		return *new(Pair[T, U]), err0
	}

	item1, err1 := z.source1.Next()
	if err1 != nil {
		return *new(Pair[T, U]), err1
	}

	result := Pair[T, U]{F0: item0, F1: item1}
	return result, nil
}
