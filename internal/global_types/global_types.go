// Package globaltypes contains the shared interfaces to be used across the application
package globaltypes

type BruteForce interface {
	GetHostAndWordlist() (host string, wordlist string)
}
