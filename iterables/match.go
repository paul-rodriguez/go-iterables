package iterables

func Match[T any](
	iter Iterable[T],
	other Iterable[T],
	equal func(T, T) bool) (bool, error) {

	zip := Zip2(iter, other)
	return zip.All(func(p Pair[T, T]) bool {
		return equal(p.F0, p.F1)
	})
}
