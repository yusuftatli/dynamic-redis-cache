package apiclient

import (
	"log"
	"net/http"
	"sync"

	"github.com/yusuftatli/hepsiburada/httpclient"
)

type Provider2Config struct {
	Url string `yaml:"url"`
}

type CurrencyProvider2 struct {
	config     Provider2Config
	httpClient httpclient.Client
}

func (p CurrencyProvider2) Handle(channel chan ProviderResponse, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	req, err := http.NewRequest("GET", p.config.Url, nil)
	if err != nil {
		log.Println("an error occured when creating request.", err)
	}

	response := ProviderResponse{}
	if err := p.httpClient.SendRequest(req, &response); err != nil {
		return
	}

	channel <- response
}
