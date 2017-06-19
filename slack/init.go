package slack

import (
	"github.com/simonjefford/fourthbot"
	"go4.org/jsonconfig"
)

// An Initializer is a function thaat takes a piece of json config and
// returns a responder or an error
type Initializer func(jsonconfig.Obj) (fourthbot.RegisteringResponder, error)

// InitializerTable is a type that stores initialization functions
// against names of responders in a config file
type InitializerTable map[string]Initializer
