package validator

import (
	"regexp"
)

// email address format check
var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// defining the validator
type Validator struct {
	Errors map[string]string
}

// new Validator that contains a map of validation errors
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid() returns true if the error map is empty
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddErrors() adds an error message to the map
func (v *Validator) AddErrors(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check() adds an error message if a validation check is not 'ok'
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddErrors(key, message)
	}
}

// In() returns true if a psecific value is in a list of strings
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// Matches() returns true if a string value matches a specific regexp pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique() returns true if all string values in a slice are unique
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
