package iterables

func (i Iterable[T]) Consume(f ...func(T)) {

	for {
		item, err := i.Next()
		if err != nil {
			return
		}
		if len(f) >= 1 {
			f[0](item)
		}
	}
}
