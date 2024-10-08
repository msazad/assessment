package cronjob

import (
	"github.com/msazad/assessment/src/service"
	"github.com/robfig/cron/v3"
)

func StartCronJob() {
	c := cron.New()
	c.AddFunc("@every 5m", func() {
		service.FetchAndStoreCryptocurrencyData()
	})
	c.Start()
}
