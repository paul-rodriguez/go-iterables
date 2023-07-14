package iterables_test

import "github.com/paul-rodriguez/go-iterables"

func TestRange(t *testing.T) {

    r := iterables.Range()
    assert.Equal(t, 0, r.Next())
}
