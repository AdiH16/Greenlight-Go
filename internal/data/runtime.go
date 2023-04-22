package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// error if parse or convert fails
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// custom Runtime type, type int32 (same as Movie struct field)
type Runtime int32

// implementation of MarshalJSON() method on Runtime type
// returns string format "<runtime> mins"
func (r Runtime) MarshalJSON() ([]byte, error) {
	// generating the string
	jsonValue := fmt.Sprintf("%d mins", r)

	// wraping in double quotes for JSON string
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

// implementation of UnmarshalJSON() method on Runtime type
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// we expect JSON value to be a string in format "<runtime> mins"
	// we remove the quotes, if not possible return error
	unqotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// split the string to isolate the part containing the number
	parts := strings.Split(unqotedJSONValue, " ")

	// make sure the parts are in expected format
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	// parse the string containing the number to int
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	// convert int32 to a Runtime and assign this to reciever
	*r = Runtime(i)

	return nil
}
