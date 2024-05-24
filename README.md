# Go State Machine

This repository contains a Go implementation of a state machine. The state machine is managed using a state manager that can handle multiple events and states simultaneously.


## Requirements

- Go version 1.18 or higher

## Installation

As a library

```bash
go get github.com/gavinin/go-state-machine.git
```

## Usage

### `NewStateManager()`

This method is used to create and initialize a new instance of `StateManager`.

StateManager supports generics, the following documents will be replaced by the `Example` class.

```go
stateManager := state.NewStateManager[Example]()
```

### `SetDuration(d time.Duration)`

This method is used to set the duration or frequency at which the `StateManager` updates its states. 

```go
stateManager.SetDuration(10 * time.Millisecond)
```

### `NewStateManagerEvent(eventType EventType, id string, events chan Test, duration time.Duration, callback CallbackFunc) *Event`

This method is used to create a new state event.

All data exchange of StateManager is completed through NewStateManagerEvent.

Whenever a new request is added, it needs to be completed through an event.

```go
event := stateManager.NewStateManagerEvent(ADD, strconv.Itoa(i), tests, duration, 
	func(events chan Example) {
	events <- Example{
	}
})
```

### `Event Type`

The following are all supported EventTypes.

```go
const (
	ADD     EventType = "ADD"
	DEL     EventType = "DEL"
	RESET   EventType = "RESET"
	PAUSE   EventType = "PAUSE"
	RUNNING EventType = "RUNNING"
)
```

- `ADD`: Add a status
- `DEL`: Delete a status
- `RESET`: Reset timer
- `PAUSE`: Pause this status
- `RUNNING`: This status is running


### `SendEvent(event Event)`

This method is used to send an event to the `StateManager`.

```go
stateManager.SendEvent(event)
```


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)