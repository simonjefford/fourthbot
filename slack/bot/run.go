package bot

import (
	"github.com/prometheus/common/log"
	"github.com/simonjefford/fourthbot/mock/responders"
	"github.com/simonjefford/fourthbot/slack"
	"github.com/simonjefford/fourthbot/slack/http"
	"go4.org/jsonconfig"
)

func Run(configFile, listenAddr string, initializers slack.InitializerTable) {
	config, err := jsonconfig.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	s := http.NewServer()

	for k, v := range config {
		log.Debugf("configuring %s", k)

		cfg := jsonconfig.Obj(v.(map[string]interface{}))
		rname := cfg.RequiredString("responder")
		// can't call validate yet
		if rname == "" {
			log.Fatal("no responder")
		}
		f, ok := initializers[rname]
		if !ok {
			log.Fatalf("unknown responder %s\n", k)
		}
		r, err := f(cfg)
		if err != nil {
			log.Fatal(err)
		}
		s.RegisterResponders(r)
	}
	s.RegisterResponder("/echo", responders.Echo)
	log.Infof("Starting server on %s", listenAddr)
	log.Fatal(s.ListenAndServe(listenAddr))
}
