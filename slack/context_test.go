package slack

import (
	"net/url"
	"testing"

	"github.com/simonjefford/fourthbot"
)

func TestCommandWithSlackData(t *testing.T) {
	tests := []struct {
		formKey  string
		value    string
		accessor func(cmd *fourthbot.Command) (string, bool)
	}{
		{formKey: "response_url", value: "http://example.com/postback", accessor: ResponseURLFromCommand},
	}

	for _, test := range tests {
		v := make(url.Values)
		v.Set(test.formKey, test.value)
		cmd, _ := fourthbot.ParseCommand("/emptycommand")
		cmd = CommandWithSlackData(cmd, v)
		s, ok := test.accessor(cmd)
		if !ok {
			t.Errorf("Expected %v to be in the context but was missing", test.value)
			continue
		}
		if s != test.value {
			t.Errorf("Expected %v to be in the context but got %+v", test.value, s)
		}
	}
}
