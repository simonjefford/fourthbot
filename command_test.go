package fourthbot

import (
	"context"
	"reflect"
	"testing"
)

func TestCommandParsing(t *testing.T) {
	c, err := ParseCommand("/echo hi")
	if err != nil {
		t.Fatal(err)
	}

	if c.Name != "echo" && !reflect.DeepEqual([]string{"hi"}, c.Args) {
		t.Fatal("The command was not parsed correctly. Got this:", c.Name, c.Args)
	}

	_, err = ParseCommand("echo hi")
	if err != ErrMissingSlash {
		t.Fatal("expected ErrMissingSlash, got", err)
	}
}

func TestCommandContext(t *testing.T) {
	c := &Command{
		Name: "/commandwithcontext",
	}
	if c.Context() != context.Background() {
		t.Errorf("Expected context.Background as default")
	}
	ctx := context.WithValue(context.Background(), "key", "value")
	c = c.WithContext(ctx)
	if c.Context() != ctx {
		t.Errorf("Expected new context after WithContext(ctx)")
	}
}
