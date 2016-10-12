package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/prometheus/common/log"
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
	buf *bytes.Buffer
}

func newSlackResponseWriter(w http.ResponseWriter) *slackResponseWriter {
	return &slackResponseWriter{
		ResponseWriter: w,
		buf:            &bytes.Buffer{},
	}
}

func (srw *slackResponseWriter) WriteStatus(s int) {
	srw.WriteHeader(s)
}

func (srw *slackResponseWriter) Write(b []byte) (int, error) {
	return srw.buf.Write(b)
}

func (srw *slackResponseWriter) WriteResponseToHTTP() {
	srw.Header().Add("Content-Type", "application/json")
	srw.buf.WriteTo(srw.ResponseWriter)
}

func (srw *slackResponseWriter) PostResponseToURL(ctx context.Context, url string) {
	r, err := http.NewRequest("POST", url, srw.buf)
	if err != nil {
		// TODO(SJJ) what do do here?
		return
	}

	r = r.WithContext(ctx)
	_, err = http.DefaultClient.Do(r)
	if err != nil {
		// TODO(SJJ) what do do here?
		return
	}
}

// A SlackServer listens for incoming /slash commands from Slack
type SlackServer struct {
	robot *fourthbot.Robot

	// maximum allowed time to return a synchronous
	// response - i.e. during the initial HTTP request
	syncResponseTimeout time.Duration

	// maximum allowed time during which POSTs can be made
	// to the response_url given by Slack
	postResponseTimeout time.Duration

	// for testing purposes
	err error
}

// Option is a func that can configure a slackserver
type Option func(s *SlackServer)

// SyncResponseTimeout returns an Option that sets the timeout period
// for a synchronous response
func SyncResponseTimeout(timeout time.Duration) Option {
	return func(s *SlackServer) {
		if timeout != 0 {
			s.syncResponseTimeout = timeout
		}
	}
}

// PostResponseTimeout return an Option that sets the timeout period
// for a posted response
func PostResponseTimeout(timeout time.Duration) Option {
	return func(s *SlackServer) {
		if timeout != 0 {
			s.postResponseTimeout = timeout
		}
	}
}

// NewServer returns a new instance of Server configured with a Robot.
func NewServer(opts ...Option) *SlackServer {
	s := &SlackServer{
		robot: fourthbot.NewRobot(),
	}

	for _, o := range opts {
		o(s)
	}

	if s.postResponseTimeout == 0 {
		s.postResponseTimeout = 30 * time.Minute
	}

	if s.syncResponseTimeout == 0 {
		s.syncResponseTimeout = 3 * time.Second
	}

	return s
}

// ListenAndServe starts up an HTTP server using s as its handler.
func (s *SlackServer) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}

// RegisterResponder registers a Responder with the underlying Robot.
func (s *SlackServer) RegisterResponder(c string, res fourthbot.Responder) {
	log.Infof("Registering a responder for %s", c)
	s.robot.RegisterResponder(c, res)
}

// RegisterResponders allows Server to use a Registrar to register one
// or more Responders.
func (s *SlackServer) RegisterResponders(res fourthbot.Registrar) {
	log.Infof("Registering multiple responders")
	res.RegisterResponders(s.robot)
}

func (s *SlackServer) handle(ctx context.Context, cmd *fourthbot.Command, srw *slackResponseWriter, finished chan bool) {
	defer func() {
		close(finished)
	}()
	err := s.robot.HandleCommand(ctx, cmd, srw)
	if err != nil {
		srw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(srw, err.Error())
		return
	}
}

func (s *SlackServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.FormValue(slackFormKeySSLCheck) != "" {
		w.WriteHeader(http.StatusOK)
		return
	}

	cmdstr := r.FormValue(slackFormKeyCommmand)
	textstr := r.FormValue(slackFormKeyText)
	if cmdstr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no command")
		return
	}

	c := &fourthbot.Command{}
	ctx := slack.ContextWithSlackData(context.Background(), r.Form)
	c.Name = cmdstr
	c.Args = strings.Split(textstr, " ")

	ctx, cancel := context.WithTimeout(ctx, s.postResponseTimeout)
	finished := make(chan bool)
	srw := newSlackResponseWriter(w)
	go s.handle(ctx, c, srw, finished)
	select {
	case <-time.After(s.syncResponseTimeout):
		log.Info("Ran out of time to provide direct response")
		fmt.Fprint(srw.ResponseWriter, "Working on it...")
		go s.waitForLongRunningOp(ctx, srw, cancel, finished)
	case <-finished:
		log.Info("Responding directly")
		cancel()
		srw.WriteResponseToHTTP()
	}
}

func (s *SlackServer) waitForLongRunningOp(ctx context.Context, srw *slackResponseWriter, cancel context.CancelFunc, finished chan bool) {
	defer cancel()
	select {
	case <-ctx.Done():
		s.err = ctx.Err()
		// op timed out
		return
	case <-finished:
		log.Info("Long running op finished")
		// can write the response to `response_url`
		url, ok := slack.ResponseURLFromContext(ctx)
		if !ok {
			// TODO(SJJ) what do do here?
			return
		}
		srw.PostResponseToURL(ctx, url)
	}
}
