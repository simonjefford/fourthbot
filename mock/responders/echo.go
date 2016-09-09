package responders

import (
	"fmt"
	"strings"

	"github.com/simonjefford/fourthbot"
)

// Echo is a test responder intended for unit tests and adhoc testing
// with the testbed.
var Echo = fourthbot.ResponderFunc(func(c *fourthbot.Command, rw fourthbot.ResponseWriter) {
	fmt.Fprintln(rw, strings.Join(c.Args, " "))
})
