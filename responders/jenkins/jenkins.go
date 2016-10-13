package jenkins

import (
	"context"
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/simonjefford/fourthbot"
	gojenk "github.com/yosida95/golang-jenkins"
	"go4.org/jsonconfig"
)

type jenkinsServer struct {
	auth     *gojenk.Auth
	addr     string
	handlers map[string]fourthbot.ResponderFunc
	client   *gojenk.Jenkins
	logger   log.Logger
	commands jsonconfig.Obj
}

func (j *jenkinsServer) configureCommand(key, def string, f fourthbot.ResponderFunc) {
	name := j.commands.OptionalString(key, def)
	j.handlers[name] = f
}

// New creates a new RegisteringResponder for interacting with a Jenkins server.
func New(cfg jsonconfig.Obj) (fourthbot.RegisteringResponder, error) {
	j := &jenkinsServer{
		logger: log.With("responder", "jenkins"),
	}

	err := j.applyConfig(cfg)
	if err != nil {
		return nil, err
	}

	j.handlers = make(map[string]fourthbot.ResponderFunc)
	j.configureCommand("jobdetails", "/jenkins-job", j.job)

	return j, nil
}

func (j *jenkinsServer) applyConfig(cfg jsonconfig.Obj) error {
	apiKey := cfg.RequiredString("apiKey")
	j.addr = cfg.RequiredString("host")
	user := cfg.RequiredString("user")
	j.commands = cfg.OptionalObject("commandMap")
	err := cfg.Validate()
	if err != nil {
		return err
	}

	j.auth = &gojenk.Auth{
		Username: user,
		ApiToken: apiKey,
	}

	return nil
}

func (j *jenkinsServer) RegisterResponders(r *fourthbot.Robot) {
	for k := range j.handlers {
		j.logger.Infof("Registering a handler for %s", k)
		r.RegisterResponder(k, j)
	}
}

func (j *jenkinsServer) job(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
	if j.client == nil {
		j.client = gojenk.NewJenkins(j.auth, j.addr)
	}

	if len(cmd.Args) == 0 {
		w.WriteStatus(500)
		fmt.Fprintln(w, "Please provide a job name")
		return
	}
	job, err := j.client.GetJob(cmd.Args[0])
	if err != nil {
		w.WriteStatus(500)
		fmt.Fprintf(w, "Error fetching job - %v\n", err)
		return
	}
	fmt.Fprintln(w, "Last successful build was", job.LastSuccessfulBuild.Url)
}

func (j *jenkinsServer) Respond(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
	h, ok := j.handlers[cmd.Name]
	if !ok {
		// TODO(SJJ) status handling
		w.WriteStatus(500)
	}

	h(ctx, cmd, w)
}
