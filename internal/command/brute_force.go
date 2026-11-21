// Package command contains the aplication commands parser and functions
package command

import (
	"flag"
	"fmt"

	globaltypes "github.com/rafael-gsr/go-path-brute-force/internal/global_types"
	"github.com/rafael-gsr/go-path-brute-force/internal/model"
)

type bruteForce struct{}

func (bf *bruteForce) usageInstructions() {
	fmt.Println("Usage: ")
	fmt.Println("-help")
	fmt.Println("-h [host]")
	fmt.Println("-w [path to wordlist]")
}

func (bf *bruteForce) ParseCommands() globaltypes.BruteForceRequest {
	misingParamsErrors := []string{}
	hostFlag := flag.String("h", "", "-h [host] -- inform host")
	helpFlag := flag.Bool("help", false, "-help -- shows help")
	wordListFlag := flag.String("w", "", "-w [wordlist] -- inform wordlist")
	cookiesListFlag := flag.String("cookies", "", "-cookies [key=value,...]")
	headersListFlag := flag.String("headers", "", "-headers [key=value,...]")
	flag.Parse()

	if *hostFlag == "" {
		misingParamsErrors = append(misingParamsErrors, "Host not informed")
		return &model.BruteForceRequest{}
	}

	if *helpFlag || len(misingParamsErrors) > 0 {
		for _, message := range misingParamsErrors {
			fmt.Println(message)
		}

		bf.usageInstructions()
		return &model.BruteForceRequest{}
	}

	request := model.NewBruteForceRequest(*hostFlag, *wordListFlag, *headersListFlag, *cookiesListFlag)
	return request
}

func NewBruteForce() globaltypes.BruteForce {
	return &bruteForce{}
}
