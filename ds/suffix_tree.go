package ds

type state interface {
	Start() int
	SetStart(start int)
	End() int
	Id() int
	Transition(char rune) state
	HasTransition(char rune) bool
	Transitions() *map[rune]state
	PutTransition(char rune, s state)
	SuffixLink() state
	SetSuffixLink(s state)
}

type mainState struct {
	id          int
	start       int
	end         int
	transitions map[rune]state
	suffixLink  state
}

func (state *mainState) Start() int {
	return state.start
}

func (state *mainState) SetStart(start int) {
	state.start = start
}

func (state *mainState) End() int {
	return state.end
}

func (state *mainState) Id() int {
	return state.id
}

func (state *mainState) Transition(char rune) state {
	return state.transitions[char]
}

func (state *mainState) HasTransition(char rune) bool {
	_, ok := state.transitions[char]
	return ok
}

func (state *mainState) Transitions() *map[rune]state {
	return &state.transitions
}

func (state *mainState) PutTransition(char rune, s state) {
	state.transitions[char] = s
}

func (state *mainState) SuffixLink() state {
	return state.suffixLink
}

func (state *mainState) SetSuffixLink(s state) {
	state.suffixLink = s
}

type auxState struct {
	*mainState
}

func (aux *auxState) transition(char rune) state {
	return aux.transitions['*']
}

func (aux *auxState) hasTransition(char rune) bool {
	return true
}

type SuffixTree struct {
	text    []rune
	root    state
	stateId int
}

func NewSuffixTree(text string) *SuffixTree {
	tree := &SuffixTree{
		// 1-based indices
		text:    []rune("*" + text + string(byte(0))),
		stateId: 0,
	}
	tree.root = tree.newState(0, 0)
	tree.root.SetSuffixLink(&auxState{
		mainState: &mainState{transitions: map[rune]state{'*': tree.root}},
	})
	s, k := tree.root, 1
	for i := 1; i < len(tree.text); i += 1 {
		s, k = tree.update(s, k, i)
	}
	return tree
}

func (tree *SuffixTree) newState(start int, end int) *mainState {
	s := &mainState{start: start, end: end, id: tree.stateId, transitions: map[rune]state{}}
	tree.stateId += 1
	return s
}

// Updates the tree with the character at index i, starting from the
// active point (s, (k, i-1)).
func (tree *SuffixTree) update(s state, k int, i int) (state, int) {
	oldr := tree.root
	t := tree.text[i]
	for {
		// Get canonical reference pair for active point
		s, k := tree.canonize(s, k, i-1)
		// Check if transition already exists, split if necessary
		endPoint, r := tree.testAndSplit(s, k, i-1, t)
		// Link up previous state r if it isn't the root
		if &oldr != &tree.root {
			oldr.SetSuffixLink(r)
		}
		// If the transition already exists, we're done
		if endPoint {
			break
		}
		// Add transition for current character t
		r.PutTransition(t, tree.newState(i, len(tree.text)-1))
		oldr = r
		// Follow suffix link
		if s.SuffixLink() != nil {
			s = s.SuffixLink()
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
func (tree *SuffixTree) canonize(s state, k int, p int) (state, int) {
	for {
		if p < k {
			break
		}
		ss := s.Transition(tree.text[k])
		if ss.End()-ss.Start() <= p-k {
			k += ss.End() - ss.Start() + 1
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
func (tree *SuffixTree) testAndSplit(s state, k int, p int, t rune) (bool, state) {
	if k <= p {
		// (1) Active point is an implicit state
		ss := s.Transition(tree.text[k])
		if t == tree.text[ss.Start()+p-k+1] {
			// (a) Next char on existing path matches current char t
			return true, s
		} else {
			// (b) Next char on existing path does not match t, need to split
			r := tree.newState(ss.Start(), ss.Start()+p-k)
			s.PutTransition(tree.text[ss.Start()], r)
			ss.SetStart(ss.Start() + p - k + 1)
			r.PutTransition(tree.text[ss.Start()], ss)
			return false, r
		}
	} else {
		// (2) Active point is an explicit state: (s, (k, p)) = (r, ε)
		if s.HasTransition(t) {
			// (a) Transition for t already exists
			return true, s
		} else {
			// (b) No transition for t yet
			return false, s
		}
	}
}

type suffixConfig struct {
	state state
	runes []rune
}

// Inefficient, only for testing correctness.
func (tree *SuffixTree) Suffixes() []string {
	stack := NewStack[suffixConfig]()
	stack.Push(suffixConfig{runes: []rune{}, state: tree.root})
	results := make([]string, 0, len(tree.text))
	for !stack.IsEmpty() {
		sc, _ := stack.Pop()
		if sc.state.Id() != 0 {
			sc.runes = append(sc.runes, tree.text[sc.state.Start():sc.state.End()+1]...)
		}
		if len(*sc.state.Transitions()) == 0 {
			// Leaf node
			results = append(results, string(sc.runes[:len(sc.runes)-1]))
		} else {
			// Non-leaf node
			for _, state := range *sc.state.Transitions() {
				new_runes := make([]rune, 0, len(sc.runes))
				copy(new_runes, sc.runes)
				stack.Push(suffixConfig{state: state, runes: new_runes})
			}
		}
	}
	return results
}
