package transport

import (
	"fmt"
	"net/http"
	"sync"
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

func (m *Multiplexer) fetch(path string) {
	var message string
	response, err := m.httpClient.GET(path)
	if err != nil {
		message = fmt.Sprintf("Error ocurred while fetching %s - %s", path, err.Error())
		fmt.Println(message)
		return
	}

	if response.StatusCode != 404 {
		message = m.formatMessage(path, response)
		fmt.Println(message)
		return
	}

	return
}

func (m *Multiplexer) Exec() {
	var wg sync.WaitGroup

	for _, path := range m.pathList {
		wg.Add(1)

		go func() {
			defer wg.Done()
			m.fetch(path)
		}()
	}

	wg.Wait()
}

func NewMultiplexer(pathList []string, httpclient HTTPClient) *Multiplexer {
	return &Multiplexer{pathList: pathList, httpClient: httpclient}
}
