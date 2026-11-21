package transport

import (
	"fmt"
	"net/http"
	"time"
)

type Multiplexer struct {
	successful []string
	failed     []string
	pathList   []string
	httpClient HTTPClient
}

func (m Multiplexer) formatMessage(path string, response *http.Response) (message string) {
	message = fmt.Sprintf("%s - %s", response.Status, response.Request.URL)
	return
}

func (m *Multiplexer) fetch(path string, c chan string) {
	var message string
	fmt.Println("path")
	response, err := m.httpClient.GET(path)
	if err != nil {
		message = fmt.Sprintf("Error ocurred while fetching %s - %s", path, err.Error())
		c <- message
		return
	}

	if response.StatusCode != 404 {
		message = m.formatMessage(path, response)
		c <- message
		return
	}

	c <- message
	return
}

func printMessage(c chan string) {
	message := <-c
	fmt.Println(message)
}

func (m *Multiplexer) Exec() {
	messageChan := make(chan string)
	for _, path := range m.pathList {
		go m.fetch(path, messageChan)
		go printMessage(messageChan)
	}

	time.Sleep(time.Second * 30)
}

func NewMultiplexer(pathList []string, httpclient HTTPClient) *Multiplexer {
	return &Multiplexer{pathList: pathList, httpClient: httpclient}
}
