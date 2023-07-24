package iterables_test

import (
	"fmt"
	"testing"

	"github.com/paul-rodriguez/go-iterables/iterables"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func TestRange(t *testing.T) {

	result, err := iterables.Range().Next()
	assert.Nil(t, err)
	assert.Equal(t, 0, result)
}

func TestApply(t *testing.T) {

	result := iterables.Apply(iterables.Range(), func(i int) (int, error) {
		return i * 2, nil
	})

	for i := 0; i < 10; i++ {
		out, err := result.Next()
		assert.Nil(t, err)
		assert.Equal(t, i*2, out)
	}
}

func TestFilter(t *testing.T) {

	result := iterables.Range().Filter(func(i int) bool {
		return i%3 == 0
	})

	for i := 0; i < 10; i++ {
		out, err := result.Next()
		assert.Nil(t, err)
		assert.Equal(t, i*3, out)
	}
}

func TestFold(t *testing.T) {

	result, err := iterables.Fold(
		iterables.Range(1, 10),
		1,
		func(s int, i int) (int, error) {
			return s * i, nil
		})
	assert.Nil(t, err)
	assert.Equal(t, result, 362880)
}

func TestZip2(t *testing.T) {

	zip := iterables.Zip2(iterables.Range(), iterables.Range(10, 20, 2))

	for i := 0; i < 5; i++ {
		pair, err := zip.Next()
		assert.Nil(t, err)
		assert.Equal(t, i, pair.F0)
		assert.Equal(t, 10+i*2, pair.F1)
	}
}

func TestTee(t *testing.T) {

	copy0, copy1 := iterables.Range(5).Tee()

	for i := 0; i < 5; i++ {
		result, err := copy0.Next()
		assert.Nil(t, err)
		assert.Equal(t, i, result)
	}

	_, err0 := copy0.Next()
	assert.ErrorIs(t, err0, iterables.IterationStop{})

	for i := 0; i < 5; i++ {
		result, err := copy1.Next()
		assert.Nil(t, err)
		assert.Equal(t, i, result)
	}

	_, err1 := copy1.Next()
	assert.ErrorIs(t, err1, iterables.IterationStop{})
}

func TestSort(t *testing.T) {

	result := iterables.Range(5, 0).Sort(func(i, j int) int {
		return i - j
	})
	expected := iterables.Range(1, 6)
	equal, err := iterables.Match(expected, result, func(i, j int) bool { return i == j })
	assert.Nil(t, err)
	assert.True(t, equal)
}

func TestToSlice(t *testing.T) {

	result, err := iterables.Range(5).ToSlice()
	assert.Nil(t, err)

	expected := []int{0, 1, 2, 3, 4}
	assert.Equal(t, expected, result)

}

func TestExtend(t *testing.T) {

	result, err := iterables.Range(5).Extend(iterables.Range(3)).ToSlice()
	assert.Nil(t, err)

	expected := []int{0, 1, 2, 3, 4, 0, 1, 2}
	assert.Equal(t, expected, result)
}

func TestTeePipeline(t *testing.T) {

	const size = 50
	tee0, tee1 := iterables.Range(size).Tee()

	result, err := iterables.Zip2(tee0.Pipeline(), tee1.Pipeline()).ToSlice()
	assert.Nil(t, err)
	expected, err := iterables.Zip2(iterables.Range(size), iterables.Range(size)).ToSlice()
	assert.Nil(t, err)

	assert.Equal(t, expected, result)
}

// Benchmarks

func BenchmarkApply(b *testing.B) {

	iter := iterables.Apply[int, int](iterables.Range(b.N), func(i int) (int, error) {
		return i / 2, nil
	})

	iter.Consume()
}

func BenchmarkSort(b *testing.B) {

	r := rand.New(rand.NewSource(1444))
	iter := iterables.FromFunction(func() (int, error) {
		return r.Int(), nil
	}).
		Take(b.N).
		Sort(func(i, j int) int {
			return j - i
		})

	iter.Consume()
}

func BenchmarkPipelineSort(b *testing.B) {

	r := rand.New(rand.NewSource(1444))
	iter := iterables.FromFunction(func() (int, error) {
		return r.Int(), nil
	}).
		Take(b.N).
		Pipeline().
		Sort(func(i, j int) int {
			return j - i
		})

	iter.Consume()
}

func BenchmarkLexSortPipeline(b *testing.B) {

	r := rand.New(rand.NewSource(1444))
	rands := iterables.FromFunction(func() (int, error) {
		return r.Int(), nil
	}).
		Take(b.N).
		Pipeline()

	strings := iterables.Apply[int, string](rands, func(i int) (string, error) {
		return fmt.Sprint(i), nil
	}).Pipeline().Sort(func(s0, s1 string) int {
		if s0 == s1 {
			return 0
		} else if s0 < s1 {
			return -1
		}
		return 1
	})

	strings.Consume()
}

func BenchmarkLexSortNoPipeline(b *testing.B) {

	r := rand.New(rand.NewSource(1444))
	rands := iterables.FromFunction(func() (int, error) {
		return r.Int(), nil
	}).
		Take(b.N)

	strings := iterables.Apply[int, string](rands, func(i int) (string, error) {
		return fmt.Sprint(i), nil
	}).Sort(func(s0, s1 string) int {
		if s0 == s1 {
			return 0
		} else if s0 < s1 {
			return -1
		}
		return 1
	})

	strings.Consume()
}

func longRunningCrunch(seed int) int {

	r := rand.New(rand.NewSource(uint64(seed)))
	randomMin, err := iterables.FromFunction(func() (int, error) {
		return r.Int(), nil
	}).
		Take(1000).
		Sort(func(i, j int) int { return i - j }).
		Next()
	if err != nil {
		message := fmt.Sprintf("Unexpected error: %s", err)
		panic(message)
	}
	return randomMin
}

func BenchmarkApplyTowerPipeline(b *testing.B) {

	r := rand.New(rand.NewSource(1444))
	inputs := iterables.FromFunction(func() (int, error) {
		return r.Int(), nil
	}).
		Take(b.N)

	levels := iterables.Range(8)
	tower, err := iterables.Fold[int, iterables.Iterable[int]](
		levels,
		inputs,
		func(iter iterables.Iterable[int], level int) (iterables.Iterable[int], error) {

			result := iterables.Apply(
				iter,
				func(input int) (int, error) {

					crunchResult := longRunningCrunch(input)
					return input ^ crunchResult, nil
				}).Pipeline()
			return result, nil
		})
	if err != nil {
		b.Fatal()
	}
	tower.Consume()
}

func BenchmarkApplyTowerNoPipeline(b *testing.B) {

	r := rand.New(rand.NewSource(1444))
	inputs := iterables.FromFunction(func() (int, error) {
		return r.Int(), nil
	}).
		Take(b.N)

	levels := iterables.Range(8)
	tower, err := iterables.Fold[int, iterables.Iterable[int]](
		levels,
		inputs,
		func(iter iterables.Iterable[int], level int) (iterables.Iterable[int], error) {

			result := iterables.Apply(
				iter,
				func(input int) (int, error) {

					crunchResult := longRunningCrunch(input)
					return input ^ crunchResult, nil
				})
			return result, nil
		})
	if err != nil {
		b.Fatal()
	}
	tower.Consume()
}
