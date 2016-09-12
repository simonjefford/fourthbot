package jenkins

import (
	"github.com/simonjefford/fourthbot"
	gojenk "github.com/yosida95/golang-jenkins"
	"go4.org/jsonconfig"
)

type jenkinsServer struct {
	auth     *gojenk.Auth
	addr     string
	handlers map[string]fourthbot.ResponderFunc
	client   *gojenk.Jenkins
}

// New creates a new Responder for interacting with a Jenkins server.
func New(cfg jsonconfig.Obj) (fourthbot.Responder, error) {
	j := &jenkinsServer{}

	err := j.applyConfig(cfg)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (j *jenkinsServer) applyConfig(cfg jsonconfig.Obj) error {
	apiKey := cfg.RequiredString("apiKey")
	j.addr = cfg.RequiredString("host")
	user := cfg.RequiredString("user")
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

func (j *jenkinsServer) Respond(cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
}
