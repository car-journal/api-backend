package cardvaultstrings

import "strings"

func Split(s, sep string) []string {
	res := make([]string, 0)
	if s == "" {
		return res
	}
	res = strings.Split(s, sep)
	return res
}
