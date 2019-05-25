package util

import "github.com/tidwall/pretty"

// Pretty pretty json output
func Pretty(json string) string {
	return string(pretty.Color(pretty.Pretty([]byte(json)), nil))
}
