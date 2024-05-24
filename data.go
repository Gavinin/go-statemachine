/**
* @Author: Gavinin
* @Date: ${DATE} ${TIME}
 */

package state

import "time"

type stateManagerData[T Type] struct {
	Name       string
	State      EventType
	Data       chan T
	Function   func(chan T)
	Duration   time.Duration
	ExpiryDate time.Time
	Times      int
}

func newStateManagerData[T Type](name string, data chan T, state EventType, fun func(chan T), duration time.Duration, times int) *stateManagerData[T] {
	return &stateManagerData[T]{
		Name:       name,
		Data:       data,
		State:      state,
		Function:   fun,
		Duration:   duration,
		ExpiryDate: time.Now().Add(duration),
		Times:      times,
	}
}
