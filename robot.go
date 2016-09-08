package fourthbot

import (
	"context"
	"errors"
	"strings"
)

// A Responder responds a command received by a Robot
type Responder interface {
	Respond(cmd *Command, rw ResponseWriter)
}

var (
	// ErrUnknownCommand is the error used when a command is
	// unknown
	ErrUnknownCommand = errors.New("Unknown command")

	// ErrMissingSlash is the error used when a command is missing
	// the leading slash
	ErrMissingSlash = errors.New("Parse error: missing slash")
)

// A ResponseWriter is used by a Responder to write response to a
// command.
type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteStatus(string)
}

// The ResponderFunc type is an adapter to allow the use of ordinary
// functions as Responders
type ResponderFunc func(cmd *Command, rw ResponseWriter)

// Respond calls f(cmd, rw)
func (f ResponderFunc) Respond(cmd *Command, rw ResponseWriter) {
	f(cmd, rw)
}

// Command represents a command received by a Robot
type Command struct {
	Name string
	Args []string
	ctx  context.Context
}

// WithContext returns a shallow copy of c with its context changed to
// ctx. The provided ctx must be non-nil.
func (c *Command) WithContext(ctx context.Context) *Command {
	if ctx == nil {
		panic("nil context")
	}
	c2 := new(Command)
	*c2 = *c
	c2.ctx = ctx
	return c2
}

// Context returns the commands's context. To change the context, use
// WithContext.
//
// The returned context is always non-nil; it defaults to the
// background context.
func (c *Command) Context() context.Context {
	if c.ctx != nil {
		return c.ctx
	}

	return context.Background()
}

// ParseCommand creates a Command from a raw string
func ParseCommand(c string) (*Command, error) {
	if !strings.HasPrefix(c, "/") {
		return nil, ErrMissingSlash
	}
	parts := strings.Split(c, " ")
	return &Command{
		Name: parts[0],
		Args: parts[1:],
	}, nil
}

// ParseCommandWithContext creates a Command from a raw string and
// additionally sets the Command's context to ctx. The provided ctx
// must be non-nil.
func ParseCommandWithContext(ctx context.Context, c string) (*Command, error) {
	if ctx == nil {
		panic("nil context")
	}
	cmd, err := ParseCommand(c)
	if err != nil {
		return nil, err
	}

	cmd.ctx = ctx

	return cmd, nil
}

// A Robot is responsible for receiving commands and dispatching them
// to the appropriate Responder
type Robot struct {
	commands map[string]Responder
	rw       ResponseWriter
}

// NewRobot initializes a new Robot
func NewRobot(rw ResponseWriter) *Robot {
	return &Robot{
		commands: make(map[string]Responder),
		rw:       rw,
	}
}

// RegisterResponder registers a Responder for a particular Command
func (r *Robot) RegisterResponder(c string, res Responder) {
	r.commands[c] = res
}

// HandleCommand dispatches a Command to the appropriate Responder
func (r *Robot) HandleCommand(c *Command) error {
	res, ok := r.commands[c.Name]
	if !ok {
		return ErrUnknownCommand
	}
	res.Respond(c, r.rw)
	return nil
}
