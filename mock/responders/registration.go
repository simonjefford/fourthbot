package responders

import (
	"github.com/simonjefford/fourthbot"
)

// RegisterAll is a convenience function for registering all the mock
// and test Responders with a Robot
func RegisterAll(r *fourthbot.Robot) {
	r.RegisterResponder("/echo", Echo)
}
