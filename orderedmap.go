package orderedmap

import (
	"bytes"
	"context"
	"fmt"
	"sort"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

// MapSlice is a type structure that acts as a map and as a slice (keeping the order of the elements).
type OrderedMap[K comparable, V any] struct {
	kv map[K]*Element[K, V] // list head and tail
	ll list[K, V]
}

func New[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		kv: make(map[K]*Element[K, V]),
	}
}

// MapItem as a string, for debugging.
func (m *OrderedMap[K, V]) String() string {
	display := "[ "

	for i := 0; i < m.Len(); i++ {
		display += m.ll.Get(i).String()

		if i+1 < m.Len() {
			display += "; "
		}
	}

	display += " ]"

	return display
}

// Get returns the value for a key. If the key does not exist, the second return
// parameter will be false and the value will be nil.
func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	var (
		val   V
		exist bool
	)

	ele, exist := m.kv[key]
	if exist {
		return ele.Value, exist
	} else {
		return val, exist
	}
}

// Set will set (or replace) a value for a key. If the key was new, then true
// will be returned. The returned value will be false if the value was replaced
// (even if the value was the same).
func (m *OrderedMap[K, V]) Set(key K, value V) bool {
	element, alreadyExist := m.kv[key]
	if alreadyExist {
		element.Value = value
		return false
	}

	element = new(Element[K, V])
	element.Key = key
	element.Value = value

	m.ll.PushElement(element)
	m.kv[key] = element
	return true
}

// GetOrDefault returns the value for a key. If the key does not exist, returns
// the default value instead.
func (m *OrderedMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if value, ok := m.kv[key]; ok {
		return value.Value
	}

	return defaultValue
}

// GetElement returns the element for a key. If the key does not exist, the
// pointer will be nil.
func (m *OrderedMap[K, V]) GetElement(key K) *Element[K, V] {
	return m.kv[key]
}

func (m *OrderedMap[K, V]) Len() int {
	return m.ll.Len()
}

//
// (De)Serializing methods:
//

// MarshalJSON for map slice.
func (m OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})

	for ele := m.ll.Front(); ele != nil; ele = ele.next {
		b, err := json.Marshal(&ele.Value)
		if err != nil {
			return nil, err
		}

		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprint(ele.Key)))
		buf.Write(b)

		if ele.next != nil {
			buf.Write([]byte{','})
		}
	}

	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

// UnmarshalJSON for map slice.
func (m *OrderedMap[K, V]) UnmarshalJSON(b []byte) error {
	if m.kv == nil {
		m.kv = make(map[K]*Element[K, V])
	}

	// temporary structure for unmarshaling
	aux := map[K]*Element[K, V]{}

	var counter int64 = 0
	ctx := context.WithValue(context.Background(), mapItemCounterType{}, &counter)

	// using default unmarshaling for maps
	if err := json.UnmarshalContext(ctx, b, &aux); err != nil {
		return errors.Wrap(err, "error while unmarshaling")
	}

	// passing the values for MapSlice
	for key, ele := range aux {
		// saving this value because it will overriden on m.ll.PushElement
		var index int64 = ele.index

		// saving the element (*Element[K,V]) to both ll and kv
		m.ll.PushElement(ele)
		m.kv[key] = ele

		// filling the other fields
		ele.index = index
		ele.Key = key
	}

	// fixing the order os the elements
	sort.Sort(&m.ll)

	// so far, so good
	return nil
}
