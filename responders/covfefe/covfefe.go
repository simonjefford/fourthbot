package covfefe

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/simonjefford/fourthbot"
	"go4.org/jsonconfig"
)

func New(cfg jsonconfig.Obj) (fourthbot.RegisteringResponder, error) {
	rand.Seed(time.Now().Unix())
	return &covfefeServer{
		tweets: []string{
			"https://twitter.com/pdccurious/status/869953309603487748",
			"https://twitter.com/Impeach_D_Trump/status/869951608355868672",
			"https://twitter.com/EW/status/869945292799520769",
			"https://twitter.com/a7_FIN_SWE/status/869944674504540161",
			"https://twitter.com/MorganTStuart/status/869941010595532801",
			"https://twitter.com/2roads_diverged/status/869937838346862592",
			"https://twitter.com/josherwalla/status/869958588856389632",
		},
	}, nil
}

type covfefeServer struct {
	tweets []string
}

func (r *covfefeServer) RegisterResponders(robot *fourthbot.Robot) {
	robot.RegisterResponder("/covfefe", r)
}

func (r *covfefeServer) Respond(ctx context.Context, cmd *fourthbot.Command, w fourthbot.ResponseWriter) {
	tweet := rand.Intn(len(r.tweets))
	fmt.Fprintf(w, "{\"response_type\": \"in_channel\",\"text\": \"%s\"}", r.tweets[tweet])
}
