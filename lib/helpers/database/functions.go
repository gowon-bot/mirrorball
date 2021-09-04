package dbhelpers

import "strings"

func EscapeForILike(term string) string {
	return strings.ReplaceAll(strings.ReplaceAll(term, "_", `\_`), "%", `\%`)
}
