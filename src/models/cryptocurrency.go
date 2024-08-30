package models

type Cryptocurrency struct {
	ID               uint   `gorm:"primart_key"`
	Name             string `gorm:"size:255"`
	Symbol           string `gorm:"size:255;unique"`
	CurrentPrice     float64
	MarketCap        float64
	Volume24h        float64
	PercentChange1h  float64
	PercentChange24h float64
	PercentChange7d  float64
}
