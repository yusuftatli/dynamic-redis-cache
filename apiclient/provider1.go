package apiclient

import (
	"log"
	"net/http"
	"sync"

	"github.com/yusuftatli/hepsiburada/httpclient"
)

type Provider1Config struct {
	Url string `yaml:"url"`
}

type CurrencyProvider1 struct {
	config     Provider1Config
	httpClient httpclient.Client
}

func (p CurrencyProvider1) Handle(channel chan ProviderResponse, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	req, err := http.NewRequest("GET", p.config.Url, nil)
	if err != nil {
		log.Println("an error occured when creating request.", err)
		return
	}

	response := ProviderResponse{}
	if err := p.httpClient.SendRequest(req, &response); err != nil {
		return
	}

	channel <- response
}
