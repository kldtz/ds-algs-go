package ds

import (
	"fmt"
	"strings"
)

type transition func(char rune) *state
type hasTransition func(char rune) bool

type state struct {
	id          int
	start       int
	end         int
	transitions map[rune]*state
	suffixLink  *state
	// Somewhat non-standard, but avoids having to implement an interface with getters and setters
	transition    transition
	hasTransition hasTransition
}

func newState(id int, start int, end int) *state {
	var s *state
	s = &state{
		id:          id,
		start:       start,
		end:         end,
		transitions: map[rune]*state{},
		transition: func(char rune) *state {
			return s.transitions[char]
		},
		hasTransition: func(char rune) bool {
			_, ok := s.transitions[char]
			return ok
		},
	}
	return s
}

func newAuxState(root *state) *state {
	var s *state
	s = &state{
		transitions: map[rune]*state{'*': root},
		transition: func(char rune) *state {
			return root
		},
		hasTransition: func(char rune) bool {
			return true
		},
	}
	return s
}

type SuffixTree struct {
	text    []rune
	root    *state
	stateId int
}

func NewSuffixTree(text string) *SuffixTree {
	tree := &SuffixTree{
		// 1-based indices
		text:    []rune("*" + text + string(byte(0))),
		stateId: 1,
	}
	tree.root = tree.newState(0, 0)
	tree.root.suffixLink = newAuxState(tree.root)
	s, k := tree.root, 1
	for i := 1; i < len(tree.text); i += 1 {
		s, k = tree.update(s, k, i)
	}
	return tree
}

func (tree *SuffixTree) newState(start int, end int) *state {
	s := newState(tree.stateId, start, end)
	tree.stateId += 1
	return s
}

// Updates the tree with the character at index i, starting from the
// active point (s, (k, i-1)).
func (tree *SuffixTree) update(s *state, k int, i int) (*state, int) {
	oldr := tree.root
	t := tree.text[i]
	for {
		// Get canonical reference pair for active point
		s, k = tree.canonize(s, k, i-1)
		// Check if transition already exists, split if necessary
		endPoint, r := tree.testAndSplit(s, k, i-1, t)
		// Link up previous state r if it isn't the root
		if oldr.id != tree.root.id {
			oldr.suffixLink = r
		}
		// If the transition already exists, we're done
		if endPoint {
			break
		}
		// Add transition for current character t
		r.transitions[t] = tree.newState(i, len(tree.text)-1)
		oldr = r
		// Follow suffix link
		if s.suffixLink != nil {
			s = s.suffixLink
		} else {
			break
		}
	}
	return s, k
}

// Turns reference pair (s, (k, p)) for state r into a canonical reference,
// that is, the closest explicit ancestor of r. If p < k, state s is
// already an explicit state and thus canonical, otherwise we follow the
// transitions to states s', as long as k' <= p.
func (tree *SuffixTree) canonize(s *state, k int, p int) (*state, int) {
	for {
		if p < k {
			break
		}
		ss := s.transition(tree.text[k])
		if ss.end-ss.start <= p-k {
			k += ss.end - ss.start + 1
			s = ss
		} else {
			break
		}
	}
	return s, k
}

// Tests if reference pair (s, (k, p)) is explicit or implicit and
// whether the transition for character t is already present.
// Thus we have four cases:
//
// (1a) Implicit and next character matches t
// (1b) Implicit and next character does not match t
// 	-> we need to split the state g(s, t_k) = s'
// (2a) Explicit and transition for t exists
// (2b) Explicit and no transition for t exists
//
// Returns existing state s or new state r (1b) and whether it is the
// endpoint, that is, a state that already has a transition t (explicitly
// or implicitly).
func (tree *SuffixTree) testAndSplit(s *state, k int, p int, t rune) (bool, *state) {
	if k <= p {
		// (1) Active point is an implicit state
		ss := s.transition(tree.text[k])
		if t == tree.text[ss.start+p-k+1] {
			// (a) Next char on existing path matches current char t
			return true, s
		} else {
			// (b) Next char on existing path does not match t, need to split
			r := tree.newState(ss.start, ss.start+p-k)
			s.transitions[tree.text[ss.start]] = r
			ss.start = ss.start + p - k + 1
			r.transitions[tree.text[ss.start]] = ss
			return false, r
		}
	} else {
		// (2) Active point is an explicit state: (s, (k, p)) = (r, ε)
		if s.hasTransition(t) {
			// (a) Transition for t already exists
			return true, s
		} else {
			// (b) No transition for t yet
			return false, s
		}
	}
}

type suffixConfig struct {
	state *state
	runes []rune
}

// Inefficient, only for testing correctness.
func (tree *SuffixTree) Suffixes() []string {
	stack := NewStack[suffixConfig]()
	stack.Push(suffixConfig{runes: []rune{}, state: tree.root})
	results := make([]string, 0, len(tree.text))
	for !stack.IsEmpty() {
		sc, _ := stack.Pop()
		if sc.state.id > 1 {
			sc.runes = append(sc.runes, tree.text[sc.state.start:sc.state.end+1]...)
		}
		if len(sc.state.transitions) == 0 {
			// Leaf node
			results = append(results, string(sc.runes[:len(sc.runes)-1]))
		} else {
			// Non-leaf node
			for _, state := range sc.state.transitions {
				new_runes := make([]rune, len(sc.runes))
				copy(new_runes, sc.runes)
				stack.Push(suffixConfig{state: state, runes: new_runes})
			}
		}
	}
	return results
}

func (tree *SuffixTree) IsSubstring(pattern string) bool {
	state := tree.root
	start, end := 0, -1
	for _, char := range pattern {
		if start <= end {
			if tree.text[start] == char {
				start += 1
			} else {
				return false
			}
		} else if state.hasTransition(char) {
			state = state.transition(char)
			start, end = state.start+1, state.end
		} else {
			return false
		}
	}
	return true
}

func (tree *SuffixTree) IsSuffix(pattern string) bool {
	return tree.IsSubstring(pattern + string(byte(0)))
}

func (tree *SuffixTree) Dot() string {
	lines := NewList[string]()
	lines.
		Append("digraph SuffixTree {").Append("graph [bgcolor=transparent];").
		Append("node [shape=circle, fixedsize=true, width=0.5];").
		Append("rankdir=LR;")
	tree.dot(tree.root.suffixLink, lines)
	lines.Append("}")
	return strings.Join(lines.ToSlice(), "\n")
}

func (tree *SuffixTree) dot(state *state, lines *List[string]) {
	if state.suffixLink != nil {
		lines.Append(fmt.Sprintf("\t%d -> %d [color=steelblue];", state.id, state.suffixLink.id))
	}
	for _, child := range state.transitions {
		text := string(tree.text[child.start : child.end+1])
		text = strings.Replace(text, string(byte(0)), "#", 1)
		lines.Append(fmt.Sprintf(
			"\t%d -> %d [label=\"%d,%d: %s\"];",
			state.id, child.id, child.start, child.end, text))
		tree.dot(child, lines)
	}
}
