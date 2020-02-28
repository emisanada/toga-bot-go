package formatter

import "strings"

// ReplaceSpaces change all spaces from a string to underscores
func ReplaceSpaces(s string) string {
	return strings.Replace(s, " ", "_", -1)
}
