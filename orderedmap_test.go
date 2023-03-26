package orderedmap

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func buildExample() *OrderedMap[string, int] {
	m := New[string, int]()
	m.Set("abc", 123)
	m.Set("def", 456)
	m.Set("ghi", 789)
	return m
}

func TestMarshal(t *testing.T) {
	m := buildExample()

	bytes, err := json.Marshal(m)
	assert.NoError(t, err)

	expected := "{\"abc\":123,\"def\":456,\"ghi\":789}"
	actual := string(bytes)

	assert.JSONEq(t, expected, actual)
}

func TestUnmarshalTyped(t *testing.T) {
	// typed example
	m := New[string, int]()

	// should be: "[ [1] => {abc 123}; [2] => {def 456}; [3] => {ghi 789} ]"
	example := []byte("{\"abc\":123,\"def\":456,\"ghi\":789}")

	// sanity check: example case is valid
	valid := json.Valid(example)
	assert.True(t, valid, "sanity check: example should be valid")

	// deserializing
	err := json.Unmarshal(example, m)
	assert.NoError(t, err)

	assert.Equal(t, m.Len(), 3)

	assert.Equal(t, 123, m.ll.Get(0).Value)
	assert.Equal(t, 456, m.ll.Get(1).Value)
	assert.Equal(t, 789, m.ll.Get(2).Value)
}

func TestUnmarshalUntyped(t *testing.T) {
	// untyped example
	m := New[string, interface{}]()

	// should be: "[ [1] => {abc 123}; [2] => {def 456}; [3] => {ghi 789} ]"
	example := []byte("{\"abc\":123,\"def\":456,\"ghi\":789}")

	// sanity check: example case is valid
	valid := json.Valid(example)
	assert.True(t, valid, "sanity check: example should be valid")

	// deserializing
	err := json.Unmarshal(example, m)
	assert.NoError(t, err)

	assert.Equal(t, m.Len(), 3)

	// TODO: is there a way to improve type inference?
}
