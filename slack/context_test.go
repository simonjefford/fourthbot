package slack

import (
	"context"
	"net/url"
	"testing"
)

func TestCommandWithSlackData(t *testing.T) {
	tests := []struct {
		formKey  string
		value    string
		accessor func(ctx context.Context) (string, bool)
	}{
		{formKey: "response_url", value: "http://example.com/postback", accessor: ResponseURLFromContext},
	}

	for _, test := range tests {
		v := make(url.Values)
		v.Set(test.formKey, test.value)
		ctx := ContextWithSlackData(context.Background(), v)
		s, ok := test.accessor(ctx)
		if !ok {
			t.Errorf("Expected %v to be in the context but was missing", test.value)
			continue
		}
		if s != test.value {
			t.Errorf("Expected %v to be in the context but got %+v", test.value, s)
		}
	}
}
