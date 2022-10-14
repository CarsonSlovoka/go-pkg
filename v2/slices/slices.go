package slices

// Index returns the index of the first occurrence of e in s,
// or -1 if not present.
func Index[E comparable](s []E, e E) int {
	for i, v := range s {
		if v == e {
			return i
		}
	}
	return -1
}

// Contains reports whether e is present in s
func Contains[E comparable](s []E, e E) bool {
	return Index(s, e) >= 0
}

// Any If one of the elements in e that are in the s then return true
func Any[E comparable](s []E, e ...E) bool {
	for _, v := range e {
		if Index(s, v) >= 0 {
			return true
		}
	}
	return false
}

// All If all the elements in e that are in the s then return true
func All[E comparable](s []E, e ...E) bool {
	if len(e) == 0 {
		return false
	}
	for _, v := range e {
		if Index(s, v) == -1 {
			return false
		}
	}
	return true
}

// ChunkBy slice分組
func ChunkBy[T any](items []T, size int) (chunks [][]T) {
	if size <= 0 {
		panic("The size must be greater than zero")
	}
	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size])
	}
	return append(chunks, items)
}
