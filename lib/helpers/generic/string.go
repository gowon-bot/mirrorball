package helpers

import "strings"

func DeepCopy(s string) string {
	var sb strings.Builder
	sb.WriteString(s)
	return sb.String()
}
