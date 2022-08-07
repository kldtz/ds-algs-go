package ds

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffixes(t *testing.T) {
	tree := NewSuffixTree("banane")
	expected := []string{
		"",
		"e",
		"ne",
		"ane",
		"nane",
		"anane",
		"banane",
	}
	actual := tree.Suffixes()
	sort.Slice(actual, func(i, j int) bool {
		return len(actual[i]) < len(actual[j])
	})

	assert.Equal(t, expected, actual)
}
