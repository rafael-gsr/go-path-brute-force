// Package globaltypes contains the shared interfaces to be used across the application
package globaltypes

type BruteForce interface {
	ParseCommands() BruteForceRequest
}

type BruteForceRequest interface {
	Host() string
	WordListPath() string
	Headers() map[string]string
	Cookies() map[string]string
}
