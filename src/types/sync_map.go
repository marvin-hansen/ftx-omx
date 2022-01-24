// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package types

import (
	"container/list"
	"sync"
)

// DRAFT RELEASE NOTES — Introduction to Go 1.18
// https://tip.golang.org/doc/go1.18

// The Go Programming Language Specification - Go 1.18 Draft (incomplete)
// https://tip.golang.org/ref/spec#Type_parameters

// Know Go: Generics [pre-order]
// https://bitfieldconsulting.com/books/generics

// Code club
// https://www.youtube.com/playlist?list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3

// Ordered Maps for Go 2 (Using Generics)
// https://medium.com/swlh/ordered-maps-for-go-using-generics-875ef3816c71

// OrderMap in Go 2.0 / 1.8 using generics with Delete Op in O(1) Time complexity
// https://www.tugberkugurlu.com/archive/implementing-ordered-map-in-go-2-0-by-using-generics-with-delete-operation-in-o-1-time-complexity
// https://gotipplay.golang.org/p/UJZsQnPRmRh

// Stores a pair of key & value of type K and V
type keyValueHolder[K comparable, V any] struct {
	key   K
	value V
}

// SyncedOrderedMap stores keys in a double linked list and keyValueHolder as List elements in a hashmap index by those keys
// All read / write operations to the map and list are mutex protected.
type SyncedOrderedMap[K comparable, V any] struct {
	sync.RWMutex
	store map[K]*list.Element
	keys  *list.List
}

// NewSyncedOrderedMap creates a new SyncedOrderedMap for keys of type K and values of type V
// To create an int, string map call:
// 	m := types.NewSyncedOrderedMap[int, string]()
func NewSyncedOrderedMap[K comparable, V any]() *SyncedOrderedMap[K, V] {
	return &SyncedOrderedMap[K, V]{
		store: make(map[K]*list.Element),
		keys:  list.New(),
	}
}

// Upsert updates the value for the given key or inserts the key value pair if the key does not exist yet.
// This is a redirect with a more precise name to the Set function which upserts by default.
func (m *SyncedOrderedMap[K, V]) Upsert(key K, val V) {
	m.Set(key, val)
}

// Set stores a pair of a key type K and a value of type V if it is not yet in the map.
// If the key is already in the map, it replaces the value in the map with the given parameter value.
func (m *SyncedOrderedMap[K, V]) Set(key K, val V) {
	m.Lock()
	defer m.Unlock()
	var e *list.Element
	if _, exists := m.store[key]; !exists {
		e = m.keys.PushBack(keyValueHolder[K, V]{
			key:   key,
			value: val,
		})
	} else {
		e = m.store[key]
		e.Value = keyValueHolder[K, V]{
			key:   key,
			value: val,
		}
	}
	m.store[key] = e
}

// Get returns a value V and true  for the given key if it exits.
// Otherwise, it returns an empty value and false.
func (m *SyncedOrderedMap[K, V]) Get(key K) (value V, ok bool) {
	m.RLock()
	defer m.RUnlock()
	val, exists := m.store[key]
	if !exists {
		return *new(V), false
	} else {
		v := val.Value.(keyValueHolder[K, V]).value
		return v, true
	}
}

// Delete removes the key from the list and the value from the map.
// If the key does not exist, nothing happen.
func (m *SyncedOrderedMap[K, V]) Delete(key K) {
	m.Lock()
	defer m.Unlock()

	e, exists := m.store[key]
	if !exists {
		return
	} else {
		m.keys.Remove(e)
		delete(m.store, key)
		return
	}
}

// Iterator returns an indexed iterator over the entire key, value collection stored in the SyncedOrderedMap.
// Example usage:
//iterator := m.Iterator()
//for {
//	i, k, v := iterator() // index, key, value
//	if i == nil {
//		break
//	}
//	fmt.Println(*k, v+" is a string")
//}
func (m *SyncedOrderedMap[K, V]) Iterator() func() (*int, *K, V) {
	m.RLock()
	defer m.RUnlock()
	e := m.keys.Front()
	j := 0
	return func() (_ *int, _ *K, _ V) {
		if e == nil {
			return
		}
		keyVal := e.Value.(keyValueHolder[K, V])
		j++
		e = e.Next()
		return func() *int { v := j - 1; return &v }(), &keyVal.key, keyVal.value
	}
}
