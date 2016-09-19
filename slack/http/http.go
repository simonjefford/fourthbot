package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/simonjefford/fourthbot"
	"github.com/simonjefford/fourthbot/slack"
)

const (
	slackFormKeyCommmand = "command"
	slackFormKeySSLCheck = "ssl_check"
	slackFormKeyText     = "text"
)

type slackResponseWriter struct {
	http.ResponseWriter
}

func (srw *slackResponseWriter) WriteStatus(s int) {
	srw.WriteHeader(s)
}

// A SlackServer listens for incoming /slash commands from Slack
type SlackServer struct {
	robot *fourthbot.Robot
}

// NewServer returns a new instance of Server configured with a Robot.
func NewServer() *SlackServer {
	return &SlackServer{fourthbot.NewRobot()}
}

// ListenAndServe starts up an HTTP server using s as its handler.
func (s *SlackServer) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}

// RegisterResponder registers a Responder with the underlying Robot.
func (s *SlackServer) RegisterResponder(c string, res fourthbot.Responder) {
	s.robot.RegisterResponder(c, res)
}

// RegisterResponders allows Server to use a Registrar to register one
// or more Responders.
func (s *SlackServer) RegisterResponders(res fourthbot.Registrar) {
	res.RegisterResponders(s.robot)
}

func (s *SlackServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.FormValue(slackFormKeySSLCheck) != "" {
		w.WriteHeader(http.StatusOK)
		return
	}

	cmdstr := r.FormValue(slackFormKeyCommmand)
	textstr := r.FormValue(slackFormKeyText)
	if cmdstr == "" {
		// TODO(SJJ) - how to handle this properly?
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	c := &fourthbot.Command{}
	ctx := slack.ContextWithSlackData(context.Background(), r.Form)
	c.Name = cmdstr
	c.Args = strings.Split(textstr, " ")

	err := s.robot.HandleCommand(ctx, c, &slackResponseWriter{w})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
}
