package main

import (
	"fmt"
	"strconv"
	"time"

	apiclient "github.com/yusuftatli/hepsiburada/apiclient"
	common "github.com/yusuftatli/hepsiburada/common"
	"github.com/yusuftatli/hepsiburada/httpclient"
)

func main(){
	//init environment
	env := common.GetEnvironment()
	
	//get config
	cfg, err := apiclient.NewDefaultConfig()
	if err != nil {
	 fmt.Errorf("failed to load config")
	}
 
	httpClient := httpclient.NewClient()
	currencyProvider := apiclient.NewCurrencyProviderHandler(cfg, httpClient)
	timeInterval, err := strconv.Atoi(common.GetEnvironment().TimeInterval)
	if err != nil {
	 timeInterval = 10
	}
	ticker := time.NewTicker(time.Duration(timeInterval) * time.Second)
	currencyProvider.GetCurrenies()
	go func() {
	 for range ticker.C {
		currencyProvider.GetCurrenies()
	 }
	}()
	 
	InitializeRoutes(env)
}