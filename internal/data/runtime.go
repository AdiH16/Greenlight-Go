package data

import (
	"fmt"
	"strconv"
)

// custom Runtime type, type int32 (sane as Movie struct field)
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
