package validations

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)


/*
   ValidateVerb check if the http verb intput is valid
   Params:
       - verb: http verb string
   Return:
       - error if the verb isn't on the list
*/
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

/*
   FormatEndpoint format the input endpoint
   if the enpoint string is `/endpoint` return without changes
   if the enpoint string is `enpoint` put the `/` and return `/endpoint`
   Put the slash before the endpoint
   Params:
       - endpoint: enpoint string
   Return:
       - string return formatted string if is necessary
*/
func FormatEndpoint(endpoint string) string {
	match, _ := regexp.MatchString(`^/`, endpoint)
	if !match {
		return fmt.Sprintf("/%s",endpoint)
	}
	return endpoint
}
