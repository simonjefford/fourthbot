package jenkins

import (
	"testing"

	"go4.org/jsonconfig"
)

func loadWithTestData(t *testing.T, path string) (*jenkinsServer, error) {
	obj, err := jsonconfig.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	r, err := New(obj)
	var j *jenkinsServer
	if err == nil {
		j = r.(*jenkinsServer)
	}

	return j, err
}

func TestNew(t *testing.T) {
	j, err := loadWithTestData(t, "testdata/testjenkins.json")
	if err != nil {
		t.Fatal(err)
	}

	if g, e := j.auth.ApiToken, "key"; g != e {
		t.Error("unexpected apiKey", g)
	}
}

func TestMissingKey(t *testing.T) {
	_, err := loadWithTestData(t, "testdata/missingkey.json")
	if err == nil {
		t.Fatal("Expected an error")
	}
}
