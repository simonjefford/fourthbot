package responders

import (
	"context"
	"fmt"
	"strings"

	"github.com/simonjefford/fourthbot"
)

// Echo is a test responder intended for unit tests and adhoc testing
// with the testbed.
var Echo = fourthbot.ResponderFunc(func(ctx context.Context, c *fourthbot.Command, rw fourthbot.ResponseWriter) {
	joined := strings.Join(c.Args, " ")
	fmt.Fprintf(rw, "{\"response_type\": \"in_channel\",\"text\": \"%s\"}", joined)
})
