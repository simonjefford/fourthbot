package rally

import "testing"

func TestConfigWithDefaultCommands(t *testing.T) {
	s, err := New(map[string]interface{}{
		"user":      "user",
		"password":  "password",
		"projectID": "projectID",
	})

	if err != nil {
		t.Fatal(err)
	}

	r := s.(*rallyServer)
	if g, e := r.user, "user"; g != e {
		t.Errorf("Unexpected user %s, expected %s", g, e)
	}
}
