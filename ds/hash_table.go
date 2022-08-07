package ds

func djb2Hash(s string) int {
	hash := 5481
	for _, c := range s {
		hash = (hash << 5) + hash + int(c)
	}
	return hash
}

type HashTableEntry[K comparable, T any] struct {
	key   K
	value T
}

type HashTable[K comparable, T any] struct {
	size         int
	loadFactor   float64
	hashFunction func(s K) int
	xs           []*HashTableEntry[K, T]
}

func NewHashTable[T any]() *HashTable[string, T] {
	return &HashTable[string, T]{
		size:         0,
		loadFactor:   0.75,
		hashFunction: djb2Hash,
		xs:           make([]*HashTableEntry[string, T], 16),
	}
}

func (table *HashTable[K, T]) Size() int {
	return table.size
}

func (table *HashTable[K, T]) rehash() {
	xs := table.xs
	table.xs = make([]*HashTableEntry[K, T], 2*cap(xs))
	for _, x := range xs {
		if x != nil {
			table.set(x.key, x.value)
		}
	}
}

// This function assumes that cap(table.xs) > table.size, otherwise it goes into an infinite loop!
func (table *HashTable[K, T]) set(key K, value T) {
	index := table.hashFunction(key) % cap(table.xs)
	for i := index; ; i = (i + 1) % cap(table.xs) {
		if table.xs[i] == nil || table.xs[i].key == key {
			table.xs[i] = &HashTableEntry[K, T]{key: key, value: value}
			break
		}
	}
}

func (table *HashTable[K, T]) Set(key K, value T) {
	if float64(table.size)/float64(cap(table.xs)) > table.loadFactor {
		table.rehash()
	}
	table.set(key, value)
	table.size += 1
}

func (table *HashTable[K, T]) Get(key K) (T, bool) {
	index := table.hashFunction(key) % cap(table.xs)
	for i := index; table.xs[i] != nil; i = (i + 1) % cap(table.xs) {
		entry := table.xs[i]
		if entry.key == key {
			return entry.value, true
		}
	}
	var zero T
	return zero, false
}
