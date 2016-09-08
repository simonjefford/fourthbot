package slack

import (
	"context"
	"net/url"

	"github.com/simonjefford/fourthbot"
)

type key int

const slackFormKey = 0

// KeysFromCommand returns the slack form values stored in the
// context of the Command.
func KeysFromCommand(cmd *fourthbot.Command) (url.Values, bool) {
	val, ok := cmd.Context().Value(slackFormKey).(url.Values)
	return val, ok
}

// CommandWithSlackKeys returns a new command with vals stored in its
// context.
func CommandWithSlackKeys(cmd *fourthbot.Command, vals url.Values) *fourthbot.Command {
	ctx := context.WithValue(cmd.Context(), slackFormKey, vals)
	return cmd.WithContext(ctx)
}
