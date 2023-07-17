package iterables_test

import (
	"testing"

	"github.com/paul-rodriguez/go-iterables/iterables"
	"github.com/stretchr/testify/assert"
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
