package slack

import (
	"context"
	"net/url"

	"github.com/simonjefford/fourthbot"
)

type key int

const slackFormKey = 0

func SlackKeysFromCommand(cmd *fourthbot.Command) (url.Values, bool) {
	val, ok := cmd.Context().Value(slackFormKey).(url.Values)
	return val, ok
}

func CommandWithSlackKeys(cmd *fourthbot.Command, vals url.Values) *fourthbot.Command {
	ctx := context.WithValue(cmd.Context(), slackFormKey, vals)
	return cmd.WithContext(ctx)
}
