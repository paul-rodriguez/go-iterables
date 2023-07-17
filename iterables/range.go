package iterables

import (
	"fmt"
)

func Range(args ...int) Iterable[int] {

	if len(args) == 0 {
		return range0()
	} else if len(args) == 1 {
		stop := args[0]
		return range1(stop)
	} else if len(args) == 2 {
		start := args[0]
		stop := args[1]
		return range2(start, stop)
	} else if len(args) == 3 {
		start := args[0]
		stop := args[1]
		step := args[2]
		return range3(start, stop, step)
	}

	invalid := invalidIterable[int]{fmt.Errorf("Invalid range: too many arguments")}
	return Iterable[int]{&invalid}
}

func range0() Iterable[int] {

	genFunc := func(state *int) (int, error) {
		result := *state
		*state++
		return result, nil
	}
	result := generator[int, int]{0, genFunc}
	return Iterable[int]{&result}
}

func range1(stop int) Iterable[int] {

	genFunc := func(state *int) (int, error) {
		if *state >= stop {
			return 0, IterationStop{}
		}
		result := *state
		*state++
		return result, nil
	}
	result := generator[int, int]{0, genFunc}
	return Iterable[int]{&result}
}

func range2(start int, stop int) Iterable[int] {

	step := 1
	if start > stop {
		step = -1
	}

	genFunc := func(state *int) (int, error) {
		if *state == stop {
			return 0, IterationStop{}
		}
		result := *state
		*state = *state + step
		return result, nil
	}
	result := generator[int, int]{start, genFunc}
	return Iterable[int]{&result}
}

func range3(start int, stop int, step int) Iterable[int] {

	genFunc := func(state *int) (int, error) {
		if step < 0 {
			if *state <= stop {
				return 0, IterationStop{}
			}
		} else {
			if *state >= stop {
				return 0, IterationStop{}
			}
		}
		result := *state
		*state = *state + step
		return result, nil
	}
	result := generator[int, int]{start, genFunc}
	return Iterable[int]{&result}
}
