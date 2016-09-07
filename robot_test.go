package fourthbot

import "testing"

func TestCommandHandling(t *testing.T) {
	c := Command("deploy")
	r := NewRobot()
	dispatched := false
	r.RegisterResponder(c, ResponderFunc(func(c Command) {
		dispatched = true
	}))
	err := r.HandleCommand(c)
	if err != nil {
		t.Fatal(err)
	}
	if !dispatched {
		t.Fatal("Command was not dispatched")
	}

	err = r.HandleCommand(Command("not registered"))
	if err != ErrUnknownCommand {
		t.Fatal("Expected ErrUnknownCommand, got", err)
	}
}
