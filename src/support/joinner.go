package support

import "strings"

func Joins(name string) string {
	s := strings.TrimSpace(name)
	url := strings.Split(s, " ")
	return strings.Join(url, "-")
}
