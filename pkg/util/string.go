package util

import (
	"strings"
)

func SubstringAfter(string string, after string) string {
	pos := strings.LastIndex(string, after)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(after)
	if adjustedPos >= len(string) {
		return ""
	}
	return string[adjustedPos:len(string)]
}

func SubstringBetween(s, start, end string) []byte {
	var between []byte
	index := strings.Index(s, start)
	if index == -1 {
		return between
	}
	index += len(start)
	for {
		char := s[index]
		if strings.HasPrefix(s[index:index+len(between)], end) {
			break
		}
		between = append(between, char)
		index++
	}
	return between
}

func StringNotEmpty(s string) bool {
	return s != ""
}
