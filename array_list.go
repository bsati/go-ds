package go_ds

// ArrayList is a port of Java's ArrayList implementation
// using an array for storing element data with dynamic resizing.
type ArrayList[T comparable] struct {
	size     int
	elements []T
}

// NewArrayList returns a new empty ArrayList with a base capacity of 10 elements
func NewArrayList[T comparable]() ArrayList[T] {
	return ArrayList[T]{
		size:     0,
		elements: make([]T, 10),
	}
}

// NewArrayListWithCapacity returns a new empty ArrayList with an element buffer
// of the given initialCapacity size.
func NewArrayListWithCapacity[T comparable](initialCapacity int) ArrayList[T] {
	return ArrayList[T]{
		size:     0,
		elements: make([]T, initialCapacity),
	}
}

// ensureCapacity resizes the underlying array if neccessary
func (l *ArrayList[T]) ensureCapacity(minCapacity int) {
	oldCapacity := len(l.elements)
	if minCapacity <= oldCapacity {
		return
	}
	newCapacity := (oldCapacity*3)/2 + 1
	if newCapacity < minCapacity {
		newCapacity = minCapacity
	}
	newElements := make([]T, newCapacity)
	copy(newElements[0:oldCapacity], l.elements)
	l.elements = newElements
}

// Add adds the given value to the list
func (l *ArrayList[T]) Add(value T) {
	l.ensureCapacity(l.size + 1)
	l.elements[l.size] = value
	l.size++
}

// Remove removes the element at the given idx if it exists
func (l *ArrayList[T]) Remove(idx int) bool {
	if idx < 0 || idx >= l.size {
		return false
	}

	numMoved := l.size - idx - 1
	if numMoved > 0 {
		copy(l.elements[idx:idx+numMoved], l.elements[idx+1:idx+numMoved+1])
	}
	l.size--
	var def T
	l.elements[l.size] = def
	return true
}

// Find searches the list for the given element.
// If the list contains the elment, Find returns the corresponding index else -1.
func (l *ArrayList[T]) Find(val T) int {
	for i, elem := range l.elements {
		if elem == val {
			return i
		}
	}
	return -1
}
