package fourthbot

import (
	"context"
	"strings"
)

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
