package validations

import "testing"

func TestValidateVerb_OkValue(t *testing.T) {
	verb := "post"
	err := ValidateVerb(verb)
	if err != nil {
		t.Errorf("Error must be nil")
	}
}

func TestValidateVerb_FailValue(t *testing.T) {
	verb := "poste"
	err := ValidateVerb(verb)
	if err == nil {
		t.Errorf("Error can't be nil")
	}
}

func TestEndpoint_OkValue(t *testing.T) {
	endpoint := "/user"
	got := FormatEndpoint(endpoint)
	if got != "/user" {
		t.Errorf("Error formatting the url: %s wants /users", got)
	}
} 
