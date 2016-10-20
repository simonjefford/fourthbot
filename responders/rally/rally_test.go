package rally

import (
	"bytes"
	"context"
	"testing"

	"github.com/simonjefford/fourthbot"
	"github.com/stretchr/testify/assert"
)

func TestConfigWithDefaultCommands(t *testing.T) {
	s, err := New(map[string]interface{}{
		"user":      "user",
		"password":  "password",
		"projectID": "projectID",
	})

	if err != nil {
		t.Fatal(err)
	}

	r := s.(*rallyServer)
	if g, e := r.user, "user"; g != e {
		t.Errorf("Unexpected user %s, expected %s", g, e)
	}

	// can't actually check the function values so just check we
	// have the expected keys
	if g := r.handlers["/new-candidate-story"]; g == nil {
		t.Errorf("addCandidateStory not registered under the expected default command name")
	}
}

func TestConfigWithCustomAndDefaultCommands(t *testing.T) {
	s, err := New(map[string]interface{}{
		"user":      "user",
		"password":  "password",
		"projectID": "projectID",
		"commandMap": map[string]interface{}{
			"newcandidatestory": "/custom-command",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	r := s.(*rallyServer)

	if g := r.handlers["/custom-command"]; g == nil {
		t.Errorf("addCandidateStory not registered under the expected custom command name")
	}

	if g := r.handlers["/rally-syntax-help"]; g == nil {
		t.Errorf("syntaxhelp not registered under the expected default command name")
	}
}

type testWriter struct {
	*bytes.Buffer
	t      *testing.T
	status int
}

func (w *testWriter) WriteStatus(s int) {
	w.t.Logf("WriteStatus(%d)", s)
	w.status = s
}

func newTestWriter(t *testing.T) *testWriter {
	w := &testWriter{
		Buffer: &bytes.Buffer{},
		t:      t,
	}

	return w
}

func TestHelp(t *testing.T) {
	s, err := New(map[string]interface{}{
		"user":      "user",
		"password":  "password",
		"projectID": "projectID",
		"commandMap": map[string]interface{}{
			"newcandidatestory": "/foo",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	cmd := &fourthbot.Command{
		Name: "/rally-syntax-help",
		Args: make([]string, 0, 0),
	}

	w := newTestWriter(t)
	s.Respond(context.Background(), cmd, w)
	t.Log(w.String())
	assert.Equal(t, 200, w.status, "unexpected status")
}
