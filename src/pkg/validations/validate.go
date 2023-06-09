package validations

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ValidateVerb(verb string) error{
	VERBS := map[string]int {
		"GET" : 0, 
		"POST" : 1, 
		"DELETE": 2, 
		"PUT": 3, 
		"PATCH": 4,
	}

	if _, exists := VERBS[strings.ToUpper(verb)]; !exists {
		return errors.New(fmt.Sprintf("%s is a invalid verb", verb))
	}

	return nil
}

func FormatEndpoint(endpoint string) string {
	match, _ := regexp.MatchString(`^/`, endpoint)
	if !match {
		return fmt.Sprintf("/%s",endpoint)
	}
	return endpoint
}
