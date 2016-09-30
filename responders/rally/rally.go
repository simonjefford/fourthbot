package rally

import (
	"context"
	"net/http"

	"github.com/simonjefford/fourthbot"
	"go4.org/jsonconfig"
)

type rallyServer struct {
	user      string
	pass      string
	projectID string
	handlers  map[string]fourthbot.ResponderFunc
	client    *http.Client
}

// New creates a new RegisteringResponder for interacting with a Jenkins server.
func New(cfg jsonconfig.Obj) (fourthbot.RegisteringResponder, error) {
	r := &rallyServer{
		client: &http.Client{},
	}
	err := r.applyConfig(cfg)
	if err != nil {
		return nil, err
	}
	r.handlers = map[string]fourthbot.ResponderFunc{
		"/new-candidate-story": r.addCandidateStory,
	}
	return r, nil
}

func (r *rallyServer) applyConfig(cfg jsonconfig.Obj) error {
	r.user = cfg.RequiredString("user")
	r.pass = cfg.RequiredString("password")
	r.projectID = cfg.RequiredString("projectID")
	return cfg.Validate()
}

func (r *rallyServer) addCandidateStory(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
}

func (r *rallyServer) RegisterResponders(robot *fourthbot.Robot) {
	for k := range r.handlers {
		robot.RegisterResponder(k, r)
	}
}

func (r *rallyServer) Respond(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
	h, ok := r.handlers[cmd.Name]
	if !ok {
		w.WriteStatus(500)
	}

	h(ctx, cmd, w)
}
