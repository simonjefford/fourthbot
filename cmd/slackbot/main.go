package main

import (
	"flag"

	"github.com/simonjefford/fourthbot"
	"github.com/simonjefford/fourthbot/responders/jenkins"
	"github.com/simonjefford/fourthbot/slack"
	"github.com/simonjefford/fourthbot/slack/run"
	"go4.org/jsonconfig"
)

var (
	configFile = flag.String("configFilePath", "config.json", "path to configuration JSON file")
	listenAddr = flag.String("listenaddr", ":8080", "listen address for http")
	initFuncs  = slack.InitializerTable{
		"jenkins": func(obj jsonconfig.Obj) (fourthbot.RegisteringResponder, error) {
			return jenkins.New(obj)
		},
	}
)

func main() {
	flag.Parse()
	run.Go(*configFile, *listenAddr, initFuncs)
}
