package utils

import "strings"

func Split(s string) []string {
	tags := strings.Split(s, ",")

	for i, tag := range tags {
		tags[i] = strings.TrimSpace(tag)
	}

	return tags
}
