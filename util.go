package kvgo

import "strings"

func FormatKey(parts ...string) string {
	return strings.Join(parts, ":")
}
