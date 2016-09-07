package fourthbot

import "errors"

// A Responder responds a command received by a Robot
type Responder interface {
	Respond(cmd Command)
}

var (
	// ErrUnknownCommand is the error used when a command is
	// unknown
	ErrUnknownCommand = errors.New("Unknown command")
)

// The ResponderFunc type is an adapter to allow the use of ordinary
// functions as Responders
type ResponderFunc func(msg Command)

// Respond calls f(msg)
func (f ResponderFunc) Respond(msg Command) {
	f(msg)
}

// Command represents a command received by a Robot
type Command string

// RegisterResponder(expression string, r Responder)
// HandleMessage(m Message)

// A Robot is responsible for receiving commands and dispatching them
// to the appropriate Responder
type Robot struct {
	commands map[Command]Responder
}

// NewRobot initializes a new Robot
func NewRobot() *Robot {
	return &Robot{
		commands: make(map[Command]Responder),
	}
}

// RegisterResponder registers a Responder for a particular Command
func (r *Robot) RegisterResponder(c Command, res Responder) {
	r.commands[c] = res
}

// HandleCommand dispatches a Command to the appropriate Responder
func (r *Robot) HandleCommand(c Command) error {
	res, ok := r.commands[c]
	if !ok {
		return ErrUnknownCommand
	}
	res.Respond(c)
	return nil
}
