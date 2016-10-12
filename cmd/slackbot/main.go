package main

import (
	"flag"

	"github.com/prometheus/common/log"
	"github.com/simonjefford/fourthbot/mock/responders"
	"github.com/simonjefford/fourthbot/responders/jenkins"
	"github.com/simonjefford/fourthbot/slack/http"
	"go4.org/jsonconfig"
)

var (
	configFile = flag.String("configFilePath", "config.json", "path to configuration JSON file")
	listenAddr = flag.String("listenaddr", ":8080", "listen address for http")
)

func configureJenkins(s *http.SlackServer, cfg jsonconfig.Obj) error {
	jenkinsCfg := cfg.RequiredObject("jenkins")
	err := cfg.Validate()
	if err != nil {
		return err
	}

	j, err := jenkins.New(jenkinsCfg)
	if err != nil {
		return err
	}
	s.RegisterResponders(j)

	return nil
}

func main() {
	flag.Parse()
	config, err := jsonconfig.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	s := http.NewServer()
	err = configureJenkins(s, config)
	if err != nil {
		log.Fatal(err)
	}
	s.RegisterResponder("/echo", responders.Echo)
	log.Infof("Starting server on %s", *listenAddr)
	log.Fatal(s.ListenAndServe(*listenAddr))
}
