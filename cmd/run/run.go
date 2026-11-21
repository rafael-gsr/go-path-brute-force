package main

import (
	"fmt"

	"github.com/rafael-gsr/go-path-brute-force/internal/command"
	"github.com/rafael-gsr/go-path-brute-force/internal/file"
	"github.com/rafael-gsr/go-path-brute-force/internal/transport"
)

func main() {
	command := command.NewBruteForce()
	request := command.ParseCommands()

	fileLines, err := file.ReadFile(request.WordListPath())
	if err != nil {
		fmt.Println("Error while reading wordlist: ", err)
		return
	}

	httpClient, err := transport.NewHTTPClient(request)
	if err != nil {
		fmt.Println("Error while creating client: ", err)
		return
	}

	multiplexer := transport.NewMultiplexer(fileLines, *httpClient)
	multiplexer.Exec()
}
