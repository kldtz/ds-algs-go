package ds

func djb2_hash(s string) int {
	hash := 5481
	for _, c := range s {
		hash = (hash << 5) + hash + int(c)
	}
	return hash
}

type HashTableEntry[T any] struct {
	key   string
	value T
}

type HashTable[T any] struct {
	size         int
	loadFactor   float64
	hashFunction func(s string) int
	xs           []*HashTableEntry[T]
}

func NewHashTable[T any]() *HashTable[T] {
	return &HashTable[T]{
		size:         0,
		loadFactor:   0.75,
		hashFunction: djb2_hash,
		xs:           make([]*HashTableEntry[T], 16),
	}
}

func (table *HashTable[T]) Size() int {
	return table.size
}

func (table *HashTable[T]) rehash() {
	xs := table.xs
	table.xs = make([]*HashTableEntry[T], 2*cap(xs))
	for _, x := range xs {
		if x != nil {
			table.set(x.key, x.value)
		}
	}
}

// This function assumes that cap(table.xs) > table.size, otherwise it goes into an infinite loop!
func (table *HashTable[T]) set(key string, value T) {
	index := table.hashFunction(key) % cap(table.xs)
	for i := index; ; i = (i + 1) % cap(table.xs) {
		if table.xs[i] == nil || table.xs[i].key == key {
			table.xs[i] = &HashTableEntry[T]{key: key, value: value}
			break
		}
	}
}

func (table *HashTable[T]) Set(key string, value T) {
	if float64(table.size)/float64(cap(table.xs)) > table.loadFactor {
		table.rehash()
	}
	table.set(key, value)
	table.size += 1
}

func (table *HashTable[T]) Get(key string) (T, bool) {
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
