package fourthbot

import (
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
