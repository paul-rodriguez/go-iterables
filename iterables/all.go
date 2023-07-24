package iterables

func (i Iterable[T]) All(predicate func(T) bool) (bool, error) {
	for {
		item, err := i.Next()
		if (err == IterationStop{}) {
			return true, nil
		} else if err != nil {
			return false, err
		} else if !predicate(item) {
			return false, nil
		}
	}
}
