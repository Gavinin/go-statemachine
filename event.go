/**
* @Author: Gavinin
* @Date: ${DATE} ${TIME}
 */

package state

import "time"

type EventType string

const (
	ADD     EventType = "ADD"
	DEL     EventType = "EDL"
	RESET   EventType = "RESET"
	PAUSE   EventType = "PAUSE"
	RUNNING EventType = "RUNNING"
)

type Event[T Type] struct {
	typ      EventType
	name     string
	data     chan T
	function func(chan T)
	duration time.Duration
	times    int
}

func (s *Event[T]) SetTimes(times int) {
	s.times = times
}
