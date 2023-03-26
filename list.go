package orderedmap

// list represents a null terminated (non-circular) intrusive doubly linked list.
// The list is immediately usable after instantiation without the need of a dedicated initialization.
type list[K comparable, V any] struct {
	root Element[K, V] // head/tail of the list (next/prev initialized with nil)
	size int           // size (initialized with 0)
}

func (l *list[K, V]) Len() int {
	return l.size
}

func (l *list[K, V]) IsEmpty() bool {
	return l.root.next == nil
}

func (l *list[K, V]) Get(i int) *Element[K, V] {
	pointer := l.root.next

	for counter := 0; counter != i; counter++ {
		if pointer == nil {
			break
		} else {
			pointer = pointer.next
		}
	}

	return pointer
}

// Front returns the first element of list l or nil if the list is empty.
func (l *list[K, V]) Front() *Element[K, V] {
	return l.root.next
}

// Back returns the last element of list l or nil if the list is empty.
func (l *list[K, V]) Back() *Element[K, V] {
	return l.root.prev
}

func (l *list[K, V]) Less(i, j int) bool {
	return l.Get(i).index < l.Get(j).index
}

func (l *list[K, V]) Swap(i, j int) {
	// space is cheap, better legibility

	iElem := l.Get(i)
	iPrevElem := iElem.prev
	iNextElem := iElem.next

	jElem := l.Get(j)
	jPrevElem := jElem.prev
	jNextElem := jElem.next

	// putting i in place
	iElem.prev = jPrevElem
	iElem.next = jNextElem

	if jPrevElem != nil {
		if jPrevElem == iElem {
			jPrevElem.next = jElem
		} else {
			jPrevElem.next = iElem
		}
	} else {
		l.root.next = iElem
	}

	if jNextElem != nil {
		if jNextElem == iElem {
			jNextElem.next = jElem
		} else {
			jNextElem.prev = iElem
		}
	} else {
		l.root.prev = iElem
	}

	// putting j in place
	jElem.prev = iPrevElem
	jElem.next = iNextElem

	if iPrevElem != nil {
		if iPrevElem == jElem {
			iPrevElem.prev = iElem
		} else {
			iPrevElem.next = jElem
		}
	} else {
		l.root.next = jElem
	}

	if iNextElem != nil {
		if iNextElem == jElem {
			iNextElem.next = iElem
		} else {
			iNextElem.prev = jElem
		}
	} else {
		l.root.prev = jElem
	}
}

// Remove removes e from its list
func (l *list[K, V]) Remove(e *Element[K, V]) {
	// keeping size for len
	l.size--

	// removing the element

	if e.prev == nil {
		l.root.next = e.next
	} else {
		e.prev.next = e.next
	}

	if e.next == nil {
		l.root.prev = e.prev
	} else {
		e.next.prev = e.prev
	}

	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
}

// PushElement inserts an element e at the back of list l.
func (l *list[K, V]) PushElement(e *Element[K, V]) {
	// keeping size for len
	l.size++

	if l.IsEmpty() {
		e.index = 1
		l.root.next = e
		l.root.prev = e
		return
	}

	e.index = l.Back().index + 1
	l.root.prev.next = e
	e.prev = l.root.prev
	l.root.prev = e
}
