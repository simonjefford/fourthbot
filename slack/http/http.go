package http

import (
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

// A Server listens for incoming /slash commands from Slack
type Server struct {
	robot *fourthbot.Robot
}

// NewServer returns a new instance of Server configured with a Robot.
func NewServer() *Server {
	return &Server{fourthbot.NewRobot()}
}

// ListenAndServe starts up an HTTP server using s as its handler.
func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}

// RegisterResponder registers a Responder with the underlying Robot.
func (s *Server) RegisterResponder(c string, res fourthbot.Responder) {
	s.robot.RegisterResponder(c, res)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	c = slack.CommandWithSlackKeys(c, r.Form)
	c.Name = cmdstr
	c.Args = strings.Split(textstr, " ")

	err := s.robot.HandleCommand(c, &slackResponseWriter{w})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
}
