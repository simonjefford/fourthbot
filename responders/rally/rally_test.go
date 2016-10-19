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

	// can't actually check the function values so just check we
	// have the expected keys
	if g := r.handlers["/new-candidate-story"]; g == nil {
		t.Errorf("addCandidateStory not registered under the expected default command name")
	}
}
