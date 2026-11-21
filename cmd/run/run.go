package main

import (
	"fmt"

	"github.com/rafael-gsr/go-path-brute-force/internal/command"
	"github.com/rafael-gsr/go-path-brute-force/internal/file"
	"github.com/rafael-gsr/go-path-brute-force/internal/transport"
)

func main() {
	command := command.NewBruteForce()
	host, wordlist := command.GetHostAndWordlist()
	fileLines, err := file.ReadFile(wordlist)
	fmt.Println("host, wordlist ", host, wordlist)
	fmt.Println("file lines ", fileLines)
	if err != nil {
		panic(err)
	}

	httpClient := transport.NewHTTPClient(host)
	multiplexer := transport.NewMultiplexer(fileLines, *httpClient)
	multiplexer.Exec()
}
