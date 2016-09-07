package responders

import (
	"fmt"

	"github.com/fourth/fourthbot"
)

type Echo struct{}

func (e Echo) Respond(c fourthbot.Command) {
	fmt.Println("echoing")
}
