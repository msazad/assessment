package repository

import (

	"github.com/msazad/assessment/src/models"
	"github.com/msazad/assessment/utils/database"
)

func GetAllCryptocurrencies() []models.Cryptocurrency {
	var cryptos []models.Cryptocurrency
	database.DB.Find(&cryptos)
	return cryptos
}

func SearchCryptocurrencies(term string) ([]models.Cryptocurrency, error) {
	var cryptocurrencies []models.Cryptocurrency
	query := "%" + term + "%"
	err := database.DB.Where("name ILIKE ? OR symbol ILIKE ?", query, query).Find(&cryptocurrencies).Error
	if err != nil {
		return nil, err
	}
	return cryptocurrencies, nil
}
