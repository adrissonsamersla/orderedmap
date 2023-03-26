package orderedmap

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/goccy/go-json"
)

type mapItemCounterType struct{}

// Element is an element of a null terminated (non circular) intrusive doubly linked list that contains the key of the correspondent element in the ordered map too.
type Element[K comparable, V any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element[K, V]

	// A counter index used for desearializing
	index int64

	// The key that corresponds to this element in the ordered map.
	Key K

	// The value stored with this element.
	Value V
}

// MapItem as a string, for debugging.
func (e *Element[K, V]) String() string {
	return fmt.Sprintf("[%d] {%v => %v}", e.index, e.Key, e.Value)
}

// Next returns the next list element or nil.
func (e *Element[K, V]) Next() *Element[K, V] {
	return e.next
}

// Prev returns the previous list element or nil.
func (e *Element[K, V]) Prev() *Element[K, V] {
	return e.prev
}

// UnmarshalJSON for orderedmap element.
func (e *Element[K, V]) UnmarshalJSON(ctx context.Context, data []byte) error {
	// storign the value to be unmarsheled
	var val V

	// using default unmarshaling for the value type
	if err := json.UnmarshalContext(ctx, data, &val); err != nil {
		return err
	}

	// retrieving the sequential counter from the context
	// it is importante that each call of UnmarshalJSON for MapSlice uses a new counter to keep order within a MapSlice
	counter := ctx.Value(mapItemCounterType{}).(*int64)

	// saving the values in this MapItem
	e.Value = val
	e.index = atomic.AddInt64(counter, 1) // using atomic here is only for precaution

	// so far, so good
	return nil
}
