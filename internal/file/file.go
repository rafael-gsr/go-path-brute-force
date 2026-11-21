// Package file contains the methods to access and parse wordlists
package file

import "os"

func ReadFile(path string) ([]string, error) {
	if path == "" {
		path = "wordlist/wordlist.txt"
	}

	rawData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := Parser(rawData)
	return lines, nil
}
