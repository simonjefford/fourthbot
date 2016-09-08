package fourthbot

import (
	"testing"

	"github.com/fourth/fourthbot/mock"
)

func TestCommandHandling(t *testing.T) {
	c, _ := ParseCommand("/deploy")
	r := NewRobot(mock.NewMockResponseWriter())
	dispatched := false
	r.RegisterResponder("/deploy", ResponderFunc(func(c *Command, rw ResponseWriter) {
		dispatched = true
	}))
	err := r.HandleCommand(c)
	if err != nil {
		t.Fatal(err)
	}
	if !dispatched {
		t.Fatal("Command was not dispatched")
	}

	c, _ = ParseCommand("/does not exist")
	err = r.HandleCommand(c)
	if err != ErrUnknownCommand {
		t.Fatal("Expected ErrUnknownCommand, got", err)
	}
}
