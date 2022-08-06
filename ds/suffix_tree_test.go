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
	sort.Strings(actual)

	assert.Equal(t, expected, tree.Suffixes())
}
