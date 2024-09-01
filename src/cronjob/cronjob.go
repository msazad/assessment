package cronjob

import (
    "github.com/robfig/cron/v3"
    "web-api/src/service"
)

func StartCronJob() {
    c := cron.New()
    c.AddFunc("@every 5m", func() {
        service.FetchAndStoreCryptocurrencyData()
    })
    c.Start()
}
