package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/adrissonsamersla/orderedmap"
)

func buildExample() *orderedmap.OrderedMap[string, int] {
	m := orderedmap.New[string, int]()
	m.Set("ihg", 123)
	m.Set("fed", 456)
	m.Set("cba", 789)
	return m
}

func main() {
	m := buildExample()

	fmt.Printf("Original data     : %s\n", m.String())

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After Marshaling  : %s\n", string(b))

	m = orderedmap.New[string, int]()
	if err := json.Unmarshal(b, m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("After Unmarshaling: %s\n", m.String())
}
