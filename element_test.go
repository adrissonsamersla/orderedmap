package orderedmap

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapElementString(t *testing.T) {
	elem := Element[string, int]{
		Key:   "abc",
		Value: 123,
		index: 1,
	}

	assert.Equal(t, "[1] {abc => 123}", elem.String())
}

func TestMapItemMarshal(t *testing.T) {
	item := Element[string, int]{}

	var counter int64 = 0
	ctx := context.WithValue(context.Background(), mapItemCounterType{}, &counter)

	err := item.UnmarshalJSON(ctx, []byte("123"))
	assert.NoError(t, err)

	assert.Equal(t, 123, item.Value)
	assert.Equal(t, int64(1), item.index)
}

func TestMapItemUnmarshal(t *testing.T) {
	item := Element[string, int]{}

	var counter int64 = 0
	ctx := context.WithValue(context.Background(), mapItemCounterType{}, &counter)

	err := item.UnmarshalJSON(ctx, []byte("123"))
	assert.NoError(t, err)

	assert.Equal(t, 123, item.Value)
	assert.Equal(t, int64(1), item.index)
}
