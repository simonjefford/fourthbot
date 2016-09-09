package fourthbot

import "errors"

// A Responder responds to a command received by a Robot
type Responder interface {
	Respond(cmd *Command, rw ResponseWriter)
}

var (
	// ErrUnknownCommand is the error used when a command is
	// unknown
	ErrUnknownCommand = errors.New("Unknown command")
)

// A ResponseWriter is used by a Responder to write response to a
// command.
type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteStatus(int)
}

// The ResponderFunc type is an adapter to allow the use of ordinary
// functions as Responders
type ResponderFunc func(cmd *Command, rw ResponseWriter)

// Respond calls f(cmd, rw)
func (f ResponderFunc) Respond(cmd *Command, rw ResponseWriter) {
	f(cmd, rw)
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
func (r *Robot) HandleCommand(c *Command, rw ResponseWriter) error {
	res, ok := r.commands[c.Name]
	if !ok {
		return ErrUnknownCommand
	}
	res.Respond(c, rw)
	return nil
}
