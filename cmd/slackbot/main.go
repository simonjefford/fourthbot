package main

import (
	"flag"

	"github.com/prometheus/common/log"
	"github.com/simonjefford/fourthbot"
	"github.com/simonjefford/fourthbot/mock/responders"
	"github.com/simonjefford/fourthbot/responders/jenkins"
	"github.com/simonjefford/fourthbot/slack/http"
	"go4.org/jsonconfig"
)

var (
	configFile = flag.String("configFilePath", "config.json", "path to configuration JSON file")
	listenAddr = flag.String("listenaddr", ":8080", "listen address for http")
	initFuncs  = map[string]func(jsonconfig.Obj) (fourthbot.RegisteringResponder, error){
		"jenkins": func(obj jsonconfig.Obj) (fourthbot.RegisteringResponder, error) {
			return jenkins.New(obj)
		},
	}
)

func main() {
	flag.Parse()
	config, err := jsonconfig.ReadFile(*configFile)
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
		f, ok := initFuncs[rname]
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
	log.Infof("Starting server on %s", *listenAddr)
	log.Fatal(s.ListenAndServe(*listenAddr))
}
