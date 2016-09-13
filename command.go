package fourthbot

import (
	"errors"
	"strings"
)

var (
	// ErrMissingSlash is the error used when a command is missing
	// the leading slash
	ErrMissingSlash = errors.New("Parse error: missing slash")
)

// Command represents a command received by a Robot
type Command struct {
	Name string
	Args []string
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
