package handlers

// import (
// 	"encoding/json"
// 	"sort"
// 	"sync"

// 	"github.com/yusuftatli/hepsiburada/apiclient"
// 	"github.com/yusuftatli/hepsiburada/common"
// 	"github.com/yusuftatli/hepsiburada/httpclient"
// 	"github.com/yusuftatli/hepsiburada/models"
// 	cache "github.com/yusuftatli/hepsiburada/redis"
// )

// type ICurrencyProvider interface {
// 	Handle(channel chan models.ProviderResponse, waitGroup *sync.WaitGroup)
// }

// type CurrencyProvider struct {
// 	cfg       common.ProviderConfig
// 	providers []ICurrencyProvider
// }

// func NewCurrencyProviderHandler(cfg *common.ProviderConfig, httpClient *httpclient.Client) *CurrencyProvider {

// 	return &CurrencyProvider{
// 		cfg: *cfg,
// 		providers: []ICurrencyProvider{
// 			apiclient.CurrencyProvider1{
// 				config:     cfg.Provider1,
// 				httpClient: *httpClient,
// 			},
// 			apiclient.CurrencyProvider2{
// 				config:     cfg.Provider2,
// 				httpClient: *httpClient,
// 			},
// 			apiclient.CurrencyProvider3{
// 				config:     cfg.Provider3,
// 				httpClient: *httpClient,
// 			},
// 		},
// 	}
// }

// //TODO: CurrencyCode enum ya da constant yapılacak.
// func (currencyProvider *CurrencyProvider) GetCurrenies() {
// 	var waitGroup sync.WaitGroup
// 	channel := make(chan models.ProviderResponse)

// 	for _, provider := range currencyProvider.providers {
// 		waitGroup.Add(1)
// 		go provider.Handle(channel, &waitGroup)
// 	}

// 	go func() {
// 		waitGroup.Wait()
// 		close(channel)
// 	}()

// 	 var arrEur []float64
// 	 var arrUsd []float64
// 	 var arrGbp []float64
// 	for msg := range channel {
// 		arrEur= append(arrEur, msg.Eur)
// 		arrUsd= append(arrUsd, msg.Usd)
// 		arrGbp= append(arrGbp, msg.Gbp)
// 	}
//  		eurModel := &models.CurrencyModel{Currency: "EURO", Data: sort.Float64Slice(arrEur)[0]}
//  		eurValue, _ := json.Marshal(eurModel)
// 		usdModel := &models.CurrencyModel{Currency: "USD", Data: sort.Float64Slice(arrUsd)[0]}
//  		usdValue, _ := json.Marshal(usdModel)
// 		gbpModel := &models.CurrencyModel{Currency: "GBP", Data: sort.Float64Slice(arrGbp)[0]}
//  		gbpValue, _ := json.Marshal(gbpModel)

// 		cache.Initialize().GetKey("EURO", eurValue)
// 		cache.Initialize().GetKey("USD", usdValue)
// 		cache.Initialize().GetKey("GBP", gbpValue)
// }