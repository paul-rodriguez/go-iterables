package iterables

func Fold[T any, U any](
	source Iterable[T],
	start U,
	f func(U, T) (U, error)) (U, error) {

	accumulator := start
	for true {
		item, err := source.Next()
		if (err == IterationStop{}) {
			break
		} else if err != nil {
			return *new(U), err
		}

		accumulator, err = f(accumulator, item)
		if err != nil {
			return *new(U), err
		}
	}
	return accumulator, nil
}
