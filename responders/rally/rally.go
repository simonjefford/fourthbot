package rally

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
	"github.com/simonjefford/fourthbot"
	"go4.org/jsonconfig"
)

type rallyServer struct {
	user       string
	pass       string
	projectID  string
	handlers   map[string]fourthbot.ResponderFunc
	client     *http.Client
	commands   jsonconfig.Obj
	commandMap map[string]string
	statusSent bool
	logger     log.Logger
}

var help = map[string]string{
	"newcandidatestory": "[name swimlane candidatesprint]",
	"help":              "no args",
}

func (r *rallyServer) configureCommand(key, def string, f fourthbot.ResponderFunc) {
	name := r.commands.OptionalString(key, def)
	r.commandMap[key] = name
	r.handlers[name] = f
}

// TODO - there's a lot of similarity here with the Jenkins
// responder. Can stuff be extracted?

// New creates a new RegisteringResponder for interacting with a Rally server.
func New(cfg jsonconfig.Obj) (fourthbot.RegisteringResponder, error) {
	r := &rallyServer{
		client: &http.Client{},
		logger: log.With("responder", "rally"),
	}
	err := r.applyConfig(cfg)
	if err != nil {
		return nil, err
	}
	r.handlers = make(map[string]fourthbot.ResponderFunc)
	r.commandMap = make(map[string]string)
	r.configureCommand("newcandidatestory", "/new-candidate-story", r.addCandidateStory)
	r.configureCommand("help", "/rally-syntax-help", r.syntaxHelp)
	return r, nil
}

func (r *rallyServer) applyConfig(cfg jsonconfig.Obj) error {
	r.user = cfg.RequiredString("user")
	r.pass = cfg.RequiredString("password")
	r.projectID = cfg.RequiredString("projectID")
	r.commands = cfg.OptionalObject("commandMap")
	return cfg.Validate()
}

func (r *rallyServer) addCandidateStory(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
}

func (r *rallyServer) syntaxHelp(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
	for cmd, name := range r.commandMap {
		fmt.Fprintf(w, "%s: %s\n", name, help[cmd])
	}
}

func (r *rallyServer) RegisterResponders(robot *fourthbot.Robot) {
	for k := range r.handlers {
		r.logger.Infof("Registering a handler for %s", k)
		robot.RegisterResponder(k, r)
	}
}

func (r *rallyServer) Respond(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
	h, ok := r.handlers[cmd.Name]
	if !ok {
		w.WriteStatus(500)
		return
	}

	h(ctx, cmd, w)
	if !r.statusSent {
		w.WriteStatus(200)
	}
}
