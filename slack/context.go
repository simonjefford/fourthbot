package slack

import (
	"context"
	"net/url"

	"github.com/simonjefford/fourthbot"
)

type key int

const (
	slackResponseURLKey = iota
)

const (
	formResponseURLKey = "response_url"
)

// ResponseURLFromCommand returns the response_url stored in the
// context of the Command.
func ResponseURLFromCommand(cmd *fourthbot.Command) (string, bool) {
	val, ok := cmd.Context().Value(slackResponseURLKey).(string)
	return val, ok
}

// CommandWithSlackData returns a new command with vals stored in its
// context.
func CommandWithSlackData(cmd *fourthbot.Command, vals url.Values) *fourthbot.Command {
	ctx := cmd.Context()
	ctx = context.WithValue(ctx, slackResponseURLKey, vals.Get(formResponseURLKey))
	return cmd.WithContext(ctx)
}
