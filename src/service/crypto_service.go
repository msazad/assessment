package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"web-api/src/models"
	"web-api/src/repository"

	"github.com/spf13/viper"
)

func FetchAndStoreCryptocurrencyData() {
	apiKey := viper.GetString("API_KEY")
	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?CMC_PRO_API_KEY=%s", apiKey)
	log.Println("Starting FetchAndStoreCryptocurrencyData at", time.Now().Format(time.RFC3339))
	defer log.Println("Completed FetchAndStoreCryptocurrencyData")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	data := result["data"].([]interface{})
	for _, item := range data {
		crypto := item.(map[string]interface{})

		repository.UpdateCryptocurrency(models.Cryptocurrency{
			Name:             crypto["name"].(string),
			Symbol:           crypto["symbol"].(string),
			CurrentPrice:     crypto["quote"].(map[string]interface{})["USD"].(map[string]interface{})["price"].(float64),
			MarketCap:        crypto["quote"].(map[string]interface{})["USD"].(map[string]interface{})["market_cap"].(float64),
			Volume24h:        crypto["quote"].(map[string]interface{})["USD"].(map[string]interface{})["volume_24h"].(float64),
			PercentChange1h:  crypto["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_1h"].(float64),
			PercentChange24h: crypto["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_24h"].(float64),
			PercentChange7d:  crypto["quote"].(map[string]interface{})["USD"].(map[string]interface{})["percent_change_7d"].(float64),
		})
	}
}
