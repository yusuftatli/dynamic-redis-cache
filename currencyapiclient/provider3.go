package currencyapiclient

import (
	"log"
	"net/http"
	"sync"

	"github.com/yusuftatli/hepsiburada/httpclient"
	"github.com/yusuftatli/hepsiburada/models"
)

type CurrencyProvider3 struct {
	config     models.Provider3Config
	httpClient httpclient.Client
}

func (p CurrencyProvider3) Handle(channel chan ProviderResponse, waitGroup *sync.WaitGroup) {
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
