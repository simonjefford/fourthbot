package main

import (
	"log"

	"github.com/simonjefford/fourthbot/mock/responders"
	"github.com/simonjefford/fourthbot/slack/http"
)

func main() {
	s := http.NewServer()
	s.RegisterResponder("/echo", responders.Echo{})
	log.Fatal(s.ListenAndServe(":8080"))
}
