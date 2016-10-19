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

func TestConfigWithCustomAndDefaultCommands(t *testing.T) {
	s, _ := New(map[string]interface{}{
		"user":      "user",
		"password":  "password",
		"projectID": "projectID",
		"commandMap": map[string]interface{}{
			"newcandidatestory": "/custom-command",
		},
	})

	r := s.(*rallyServer)

	if g := r.handlers["/custom-command"]; g == nil {
		t.Errorf("addCandidateStory not registered under the expected custom command name")
	}

	if g := r.handlers["/rally-syntax-help"]; g == nil {
		t.Errorf("syntaxhelp not registered under the expected default command name")
	}
}
