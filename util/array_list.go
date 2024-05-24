/**
* @Author: Gavinin
* @Date: ${DATE} ${TIME}
 */

package util

import (
	"sync"
)

type SafeListType interface {
	string | int | interface{}
}

type SafeList[T SafeListType] struct {
	zero T
	mu   sync.RWMutex
	list []T
}

func NewSafeList[T SafeListType]() SafeList[T] {
	return SafeList[T]{
		list: make([]T, 0),
	}
}

func (s *SafeList[T]) Append(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.list = append(s.list, item)
}

func (s *SafeList[T]) Get(index int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if index >= len(s.list) {
		return s.zero, false
	}
	return s.list[index], true
}

func (s *SafeList[T]) Remove(index int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if index >= len(s.list) {
		return false
	}
	s.list = append(s.list[:index], s.list[index+1:]...)
	return true
}

func (s *SafeList[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.list)
}

func (s *SafeList[T]) AsList() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.list
}
