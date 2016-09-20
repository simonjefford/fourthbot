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
	status           int
	responderContext context.Context
}

func (r *testResponder) Respond(ctx context.Context, cmd *fourthbot.Command, rw fourthbot.ResponseWriter) {
	r.called = true
	r.responderContext = ctx
	rw.WriteStatus(r.status)
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

func runTestWithRegistrar(t *testing.T, reg *registrar, cmdstr string, form url.Values) *httptest.ResponseRecorder {
	s := NewServer()
	s.RegisterResponders(reg)
	form.Set("command", cmdstr)
	r := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	return w
}

func TestResponseFromResponder(t *testing.T) {
	tr := &testResponder{
		name:   "fooResponder",
		status: 200,
	}
	w := runTest(t, tr, "/foo", map[string][]string{})
	if w.Code != 200 {
		t.Errorf("Expected 200 got %d (\"%s\" written to the response)", w.Code, w.Body.String())
	}
	if !tr.called {
		t.Error("responder not called")
	}
}

func TestResponderStatusTreatedAsHTTPStatus(t *testing.T) {
	tr := &testResponder{
		name:   "fooResponder",
		status: 400,
	}
	w := runTest(t, tr, "/foo", map[string][]string{})
	if w.Code != 400 {
		t.Errorf("Expected 400 got %d", w.Code)
	}
}

// TODO(SJJ) - this test is probably in the wrong place...
func TestContextPropagation(t *testing.T) {
	tr := &testResponder{
		name: "fooResponder",
	}
	hook := "https://slack.com/hook"
	runTest(t, tr, "/foo", map[string][]string{
		"response_url": []string{hook},
	})
	if g, e := tr.responderContext.Value(0), hook; g != e {
		t.Errorf("Expected %s got %v", e, g)
	}
}

type registrar struct {
	res *testResponder
}

func (r *registrar) RegisterResponders(robot *fourthbot.Robot) {
	robot.RegisterResponder("/foo", r.res)
}

func TestRegistrarHandling(t *testing.T) {
	tr := &testResponder{
		name: "fooResponder",
	}
	r := &registrar{tr}
	runTestWithRegistrar(t, r, "/foo", map[string][]string{})
	if !tr.called {
		t.Errorf("Expected responder to be called but was not")
	}
}

func TestSSLCheck(t *testing.T) {
	s := NewServer()
	r := httptest.NewRequest("POST", "/?ssl_check=1", strings.NewReader(make(url.Values).Encode()))

	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Errorf("Expected 200 on an ssl_check, got %d. Body was %s", w.Code, w.Body.String())
	}
}
