package ds

func djb2_hash(s string) int {
	hash := 5481
	for _, c := range s {
		hash = (hash << 5) + hash + int(c)
	}
	return hash
}

type HashTable[T any] struct {
	size         int
	loadFactor   float64
	hashFunction func(s string) int
	xs           []*T
}

func NewHashTable[T any](initialCapacity int, loadFactor float64, hashFunction func(s string) int) *HashTable[T] {
	if initialCapacity == 0 {
		initialCapacity = 16
	}
	if loadFactor == 0 {
		loadFactor = 0.75
	}
	if hashFunction == nil {
		hashFunction = djb2_hash
	}
	return &HashTable[T]{
		size:         0,
		loadFactor:   loadFactor,
		hashFunction: hashFunction,
		xs:           make([]*T, 0, initialCapacity),
	}
}

// func (table *HashTable[T]) rehash() *HashTable[T] {
// 	new_xs := make([]T, 0, 2*cap(table.xs))
// 	for _, x := range table.xs {
// 		if x != nil {

// 		}
// 	}
// }

func (table *HashTable[T]) Set(key string, value T) {
	if float64(table.size)/float64(cap(table.xs)) > table.loadFactor {
		// TODO: rehash
	}
	index := table.hashFunction(key) % cap(table.xs)
	for i := index; ; i = (i + 1) % cap(table.xs) {
		if table.xs[i] == nil {
			table.xs[i] = &value
			break
		}
	}
	table.size += 1
}
