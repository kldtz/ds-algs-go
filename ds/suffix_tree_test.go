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

func TestIsSubstring(t *testing.T) {
	text := "banane"
	tree := NewSuffixTree(text)
	substrings := make([]string, len(text)*(len(text)+1)/2)
	for i := 0; i < len(text); i += 1 {
		for j := i + 1; j <= len(text); j += 1 {
			substrings = append(substrings, text[i:j])
		}
	}

	for _, substring := range substrings {
		assert.True(t, tree.IsSubstring(substring))
	}
	assert.True(t, tree.IsSubstring(""))
	assert.False(t, tree.IsSubstring("xyz"))
	assert.False(t, tree.IsSubstring("bna"))
	assert.False(t, tree.IsSubstring(" "))
}

func TestIsSuffix(t *testing.T) {
	text := "lagerregal"
	tree := NewSuffixTree(text)

	for i := 0; i <= len(text); i += 1 {
		assert.True(t, tree.IsSuffix(text[i:]))
	}

	assert.False(t, tree.IsSuffix("lager"))
	assert.False(t, tree.IsSuffix("g"))
}

func TestDot(t *testing.T) {
	text := "abaabcab"
	tree := NewSuffixTree(text)

	dot := tree.Dot()
	t.Logf(dot)
}
