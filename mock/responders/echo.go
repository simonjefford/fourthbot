package responders

import (
	"fmt"
	"strings"

	"github.com/simonjefford/fourthbot"
)

// Echo is a test type intended for unit tests and adhoc testing with
// the testbed.
type Echo struct{}

// Respond provides the implementation for Responder
func (e Echo) Respond(c *fourthbot.Command, rw fourthbot.ResponseWriter) {
	fmt.Fprintln(rw, strings.Join(c.Args, " "))
}
