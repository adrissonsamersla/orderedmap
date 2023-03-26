package orderedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListLen(t *testing.T) {
	var ll list[string, int]
	assert.Zero(t, ll.Len())

	var ele Element[string, int]

	ll.PushElement(&ele)
	assert.Equal(t, 1, ll.Len())

	ll.Remove(&ele)
	assert.Zero(t, ll.Len())
}

func TestListSwap(t *testing.T) {
	var ll list[string, int]

	elem0 := Element[string, int]{Key: "a", Value: 0}
	elem1 := Element[string, int]{Key: "b", Value: 1}

	ll.PushElement(&elem0)
	ll.PushElement(&elem1)

	assert.Equal(t, "a", ll.Get(0).Key)
	assert.Equal(t, "b", ll.Get(1).Key)

	ll.Swap(0, 1)

	assert.Equal(t, "b", ll.Get(0).Key)
	assert.Equal(t, "a", ll.Get(1).Key)

}
