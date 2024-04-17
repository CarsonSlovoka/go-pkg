package set

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"sync"
)

type MapSet[T comparable] struct {
	elem map[T]struct{}
	sync.RWMutex
}

func NewMapSet[T comparable](elems ...T) *MapSet[T] {
	s := MapSet[T]{}
	s.elem = make(map[T]struct{})
	s.Add(elems...)
	return &s
}

func (s *MapSet[T]) Add(keys ...T) {
	s.Lock()
	defer s.Unlock()
	for _, key := range keys {
		if _, exists := s.elem[key]; exists {
			continue
		}
		s.elem[key] = struct{}{}
	}
}

func (s *MapSet[T]) Delete(keys ...T) {
	s.Lock()
	defer s.Unlock()
	for _, key := range keys {
		if _, exists := s.elem[key]; !exists {
			continue
		}
		delete(s.elem, key)
	}
}

func (s *MapSet[T]) Clear(keys ...T) {
	s.Lock()
	defer s.Unlock()
	clear(s.elem)
}

func (s *MapSet[T]) Has(key T) bool {
	s.RLock()
	defer s.RUnlock()
	_, exists := s.elem[key]
	return exists
}

func (s *MapSet[T]) Contains(keys ...T) bool {
	for _, key := range keys {
		if !s.Has(key) {
			return false
		}
	}
	return true
}

func (s *MapSet[_]) Size() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.elem)
}

func (s *MapSet[_]) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.elem) == 0
}

func (s *MapSet[T]) Items() []T {
	s.RLock()
	defer s.RUnlock()
	var items []T
	for key, _ := range s.elem {
		items = append(items, key)
	}
	return items
}

func (s *MapSet[T]) Union(s2 *MapSet[T]) *MapSet[T] {
	s.RLock()
	s2.RLock()
	defer func() {
		s.RUnlock()
		s2.RUnlock()
	}()

	all := NewMapSet[T](s.Items()...)
	all.Add(s2.Items()...)
	return all
}

func (s *MapSet[T]) Diff(s2 *MapSet[T]) *MapSet[T] {
	s.RLock()
	s2.RLock()
	defer func() {
		s.RUnlock()
		s2.RUnlock()
	}()

	all := s.Union(s2)

	for _, k := range all.Items() {
		if (s.Has(k) && s2.Has(k)) ||
			(s2.Has(k) && s.Has(k)) {
			all.Delete(k)
		}
	}
	return all
}

func (s *MapSet[T]) String() string {
	s.RLock()
	defer s.RUnlock()
	var strs []string
	for key, _ := range s.elem { // 因為map是無序的，所以取出來的順序可能不同
		strs = append(strs, fmt.Sprintf("%v", key))
	}

	// make sure the order do not change
	slices.SortFunc(strs, func(a, b string) int {
		return cmp.Compare(a, b)
	})
	return strings.Join(strs, " ")
}

func (s *MapSet[T]) Clone() *MapSet[T] {
	s.RLock()
	defer s.RUnlock()
	s2 := NewMapSet[T]()
	var ks []T
	for key, _ := range s.elem {
		ks = append(ks, key)
	}
	s2.Add(ks...)
	return s2
}
