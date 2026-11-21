// Package model implements the common dtos for the application
package model

import (
	"fmt"
	"strings"
)

type BruteForceRequest struct {
	host         string
	wordlistPath string
	headers      map[string]string
	cookies      map[string]string
}

func (bfr *BruteForceRequest) Host() string               { return bfr.host }
func (bfr *BruteForceRequest) WordListPath() string       { return bfr.wordlistPath }
func (bfr *BruteForceRequest) Headers() map[string]string { return bfr.headers }
func (bfr *BruteForceRequest) Cookies() map[string]string { return bfr.cookies }

func parseStringToMap(input string) map[string]string {
	mappedValues := make(map[string]string)
	keysAndValues := strings.Split(input, ",")

	for _, pair := range keysAndValues {
		keyValuePair := strings.Split(pair, "=")

		if len(keyValuePair) < 2 {
			fmt.Sprintf("warning: incomplete value cannot be parsed: %s", pair)
			continue
		}

		mappedValues[keyValuePair[0]] = keyValuePair[1]
	}

	return mappedValues
}

func NewBruteForceRequest(host string, wordlistPath string, headers string, cookies string) *BruteForceRequest {
	mappedHeaders := parseStringToMap(headers)
	mappedCookies := parseStringToMap(cookies)

	return &BruteForceRequest{host, wordlistPath, mappedHeaders, mappedCookies}
}
