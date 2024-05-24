/**
* @Author: Gavinin
* @Date: ${DATE} ${TIME}
 */

package state

import (
	"errors"
	"fmt"
	"github.com/gavinin/go-state-machine/util"
	"time"
)

type Type interface {
	chan bool | string | interface{}
}

type Manager[T Type] struct {
	zero                 stateManagerData[T]
	registeredStatusList util.SafeList[stateManagerData[T]]
	duration             time.Duration
}

// NewStateManager This method is used to create and initialize a new instance of StateManager.
func NewStateManager[T Type]() *Manager[T] {
	s := &Manager[T]{
		registeredStatusList: util.NewSafeList[stateManagerData[T]](),
	}
	go func() {
		for {
			for _, t := range s.registeredStatusList.AsList() {
				if t.State == PAUSE {
					continue
				}
				now := time.Now()
				if now.After(t.ExpiryDate) {
					ts := t
					if t.Times == 0 {
						s.removeState(t.Name)
					} else if t.Times > 0 {
						ts.Times -= 1
						s.setRegisteredStatus(ts)
					}
					ts.Function(t.Data)
					s.resetTimer(t.Name)
					if now.Sub(t.ExpiryDate) > time.Second {
						fmt.Println("out of date")
					}
				}
			}
			time.Sleep(s.duration)
		}
	}()

	return s
}

// SetDuration This method is used to set the duration or frequency at which the StateManager updates its states.
func (t *Manager[T]) SetDuration(duration time.Duration) {
	t.duration = duration
}

// NewStateManagerEvent This method is used to create a new state event.
func (t *Manager[T]) NewStateManagerEvent(typ EventType, name string, data chan T, duration time.Duration, function func(chan T)) Event[T] {
	return Event[T]{typ: typ, name: name, data: data, duration: duration, function: function, times: -1}
}

// SendEvent This method is used to send an event to the StateManager.
func (t *Manager[T]) SendEvent(event Event[T]) error {
	if event.name == "" {
		return errors.New("name should not be empty")
	}
	switch event.typ {
	case ADD:
		if t.hasExist(event.name) {
			return errors.New("name has existed")
		}
		data := newStateManagerData(event.name, event.data, RUNNING, event.function, event.duration, event.times)
		t.registeredStatusList.Append(*data)
		break
	case DEL:
		if !t.removeState(event.name) {
			return errors.New("name not found")
		}
		break
	case RESET:
		if !t.resetTimer(event.name) {
			return errors.New("name not found")
		}
		break
	case PAUSE:
		if !t.changeState(event.name, PAUSE) {
			return errors.New("name not found")
		}
	case RUNNING:
		if !t.changeState(event.name, RUNNING) {
			return errors.New("name not found")
		}

	}
	return nil
}

func (t *Manager[T]) getByName(name string) (stateManagerData[T], bool) {
	for _, s := range t.registeredStatusList.AsList() {
		if s.Name == name {
			return s, true
		}
	}
	return t.zero, false
}

func (t *Manager[T]) hasExist(name string) bool {
	_, b := t.getByName(name)
	return b
}

func (t *Manager[T]) removeState(name string) bool {
	for i, s := range t.registeredStatusList.AsList() {
		if s.Name == name {
			t.registeredStatusList.Remove(i)
			return true
		}
	}
	return false
}

func (t *Manager[T]) changeState(name string, status EventType) bool {
	ts, b := t.getByName(name)
	if b {
		ts.State = status
		t.setRegisteredStatus(ts)
		return true
	}
	return false
}

func (t *Manager[T]) setExpire(name string, duration time.Duration) bool {
	ts, b := t.getByName(name)
	if b {
		ts.ExpiryDate = time.Now().Add(duration)
		t.setRegisteredStatus(ts)
		return true
	}

	return false
}

func (t *Manager[T]) resetTimer(name string) bool {
	byName, b := t.getByName(name)
	if b {
		return t.setExpire(name, byName.Duration)
	}
	return false
}

func (t *Manager[T]) setRegisteredStatus(ts stateManagerData[T]) bool {
	for i, s := range t.registeredStatusList.AsList() {
		if s.Name == ts.Name {
			t.registeredStatusList.Remove(i)
			t.registeredStatusList.Append(ts)
			return true
		}
	}
	return false
}
