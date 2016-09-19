package http

import (
	"context"
	"fmt"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/simonjefford/fourthbot"
)

type testResponder struct {
	name             string
	called           bool
	responderContext context.Context
}

func (r *testResponder) Respond(ctx context.Context, cmd *fourthbot.Command, rw fourthbot.ResponseWriter) {
	r.called = true
	r.responderContext = ctx
	fmt.Fprintf(rw, "called by %s", r.name)
}

func runTest(t *testing.T, tr *testResponder, cmdstr string, form url.Values) *httptest.ResponseRecorder {
	s := NewServer()
	s.RegisterResponder(cmdstr, tr)
	form.Set("command", cmdstr)
	r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	return w
}

func TestResponseFromResponder(t *testing.T) {
	tr := &testResponder{
		name: "fooResponder",
	}
	w := runTest(t, tr, "/foo", map[string][]string{})
	if w.Code != 200 {
		t.Errorf("Expected 200 got %d (\"%s\" written to the response)", w.Code, w.Body.String())
	}
	if !tr.called {
		t.Error("responder not called")
	}
}
