package fourthbot

import (
	"errors"
	"strings"
)

// A Responder responds a command received by a Robot
type Responder interface {
	Respond(cmd *Command)
}

var (
	// ErrUnknownCommand is the error used when a command is
	// unknown
	ErrUnknownCommand = errors.New("Unknown command")

	// ErrMissingSlash is the error used when a command is missing
	// the leading slash
	ErrMissingSlash = errors.New("Parse error: missing slash")
)

// The ResponderFunc type is an adapter to allow the use of ordinary
// functions as Responders
type ResponderFunc func(cmd *Command)

// Respond calls f(msg)
func (f ResponderFunc) Respond(cmd *Command) {
	f(cmd)
}

// Command represents a command received by a Robot
type Command struct {
	name string
	args []string
}

// ParseCommand creates a Command from a raw string
func ParseCommand(c string) (*Command, error) {
	if !strings.HasPrefix(c, "/") {
		return nil, ErrMissingSlash
	}
	parts := strings.Split(c, " ")
	return &Command{
		name: parts[0],
		args: parts[1:],
	}, nil
}

// A Robot is responsible for receiving commands and dispatching them
// to the appropriate Responder
type Robot struct {
	commands map[string]Responder
}

// NewRobot initializes a new Robot
func NewRobot() *Robot {
	return &Robot{
		commands: make(map[string]Responder),
	}
}

// RegisterResponder registers a Responder for a particular Command
func (r *Robot) RegisterResponder(c string, res Responder) {
	r.commands[c] = res
}

// HandleCommand dispatches a Command to the appropriate Responder
func (r *Robot) HandleCommand(c *Command) error {
	res, ok := r.commands[c.name]
	if !ok {
		return ErrUnknownCommand
	}
	res.Respond(c)
	return nil
}
