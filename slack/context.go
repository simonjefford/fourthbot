package slack

import (
	"context"
	"net/url"
)

type key int

const (
	slackResponseURLKey = iota
)

const (
	formResponseURLKey = "response_url"
)

// ResponseURLFromContext returns the response_url stored in the
// ctx.
func ResponseURLFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(slackResponseURLKey).(string)
	return val, ok
}

// ContextWithSlackData returns a new command with vals stored in its
// context.
func ContextWithSlackData(ctx context.Context, vals url.Values) context.Context {
	return context.WithValue(ctx, slackResponseURLKey, vals.Get(formResponseURLKey))
}
