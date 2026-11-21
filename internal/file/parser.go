package file

import (
	"strings"
)

func Parser(rawData []byte) []string {
	readableContent := string(rawData)
	contentLines := strings.Split(readableContent, "\n")
	return contentLines
}
