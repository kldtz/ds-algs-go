package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLcsNaive(t *testing.T) {
	lcs := LcsNaive("lagerregale", "illegal")
	assert.Equal(t, []string{"egal"}, lcs)

	lcs = LcsNaive("abc", "xyz")
	assert.Equal(t, []string{}, lcs)
}

func TestLcsDynamicProgramming(t *testing.T) {
	lcs := LcsDynamicProgramming("lagerregale", "illegal")
	assert.Equal(t, []string{"egal"}, lcs)

	lcs = LcsDynamicProgramming("xbcdxefgx", "ybcdyefgy")
	assert.Equal(t, []string{"bcd", "efg"}, lcs)

	lcs = LcsDynamicProgramming("abc", "xyz")
	assert.Equal(t, []string{}, lcs)
}
