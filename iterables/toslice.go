package iterables

func (i Iterable[T]) ToSlice() ([]T, error) {

	var result []T
	for {
		item, err := i.Next()
		if (err == IterationStop{}) {
			return result, nil
		} else if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
}
