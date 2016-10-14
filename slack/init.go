package slack

import (
	"github.com/simonjefford/fourthbot"
	"go4.org/jsonconfig"
)

// InitializerTable is a type that stores initialization functions
// against names of responders in a config file
type InitializerTable map[string]func(jsonconfig.Obj) (fourthbot.RegisteringResponder, error)
