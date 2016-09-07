package responders

import (
	"github.com/fourth/fourthbot"
)

func RegisterAll(r *fourthbot.Robot) {
	r.RegisterResponder(fourthbot.Command("/echo"), Echo{})
}
