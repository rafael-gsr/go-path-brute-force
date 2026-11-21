// Package command contains the aplication commands parser and functions
package command

import (
	"flag"
	"fmt"

	globaltypes "github.com/rafael-gsr/go-path-brute-force/internal/global_types"
)

type bruteForce struct{}

func (bf *bruteForce) usageInstructions() {
	fmt.Println("Usage: ")
	fmt.Println("-help")
	fmt.Println("-h [host]")
	fmt.Println("-w [path to wordlist]")
}

func (bf *bruteForce) GetHostAndWordlist() (host string, wordlist string) {
	misingParamsErrors := []string{}
	hostFlag := flag.String("h", "", "-h [host] -- inform host")
	helpFlag := flag.Bool("help", false, "-help -- shows help")
	wordListFlag := flag.String("w", "", "-w [wordlist] -- inform wordlist")
	flag.Parse()

	if *hostFlag == "" {
		misingParamsErrors = append(misingParamsErrors, "Host not informed")
		return "", ""
	}

	if *helpFlag || len(misingParamsErrors) > 0 {
		for _, message := range misingParamsErrors {
			fmt.Println(message)
		}

		bf.usageInstructions()
		return "", ""
	}

	return *hostFlag, *wordListFlag
}

func NewBruteForce() globaltypes.BruteForce {
	return &bruteForce{}
}
