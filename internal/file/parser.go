package file

import (
	"fmt"
	"strings"
)

func Parser(rawData []byte) []string {
	readableContent := string(rawData)
	contentLines := strings.Split(readableContent, "\n")
	fmt.Println("parser: ", contentLines)
	return contentLines
}
