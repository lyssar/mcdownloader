package utils

import (
	"strings"
)

func ReplaceNamed(format string, data map[string]string) string {
	var replace []string
	for key, element := range data {
		replace = append(replace, "{"+key+"}", element)
	}
	r := strings.NewReplacer(replace...)
	return r.Replace(format)
}
