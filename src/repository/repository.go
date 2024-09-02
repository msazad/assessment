package repository

import (
	"log"
	"web-api/src/models"
	"web-api/utils/database"
)

var Repo PostgresRepository

type PostgresRepositoryRepo struct{}

func PostgresInit() {
	Repo = &PostgresRepositoryRepo{}
}

/*** Inserting data to the database ***/
func (r *PostgresRepositoryRepo) Insert(req interface{}) error {
	if err := database.DB.Debug().Create(req).Error; err != nil {
		return err
	}
	return nil
}

/*** Fetching data from the database ***/
func (r *PostgresRepositoryRepo) FindById(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("unique_id = ? ", id).Find(obj).Error; err != nil {
		return err
	}
	return nil
}

/*** Updating data in the database ***/
func (r *PostgresRepositoryRepo) Update(obj interface{}, id int, update interface{}) error {
	if err := database.DB.Debug().Where("unique_id IN (?) ", id).First(obj).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

/*** Deleting data from the database ***/
func (r *PostgresRepositoryRepo) Delete(obj interface{}, id int) error {
	if err := database.DB.Debug().Where("unique_id IN (?) ", id).First(obj).Delete(obj).Error; err != nil {
		return err
	}
	return nil
} // UpdateCryptocurrency updates the cryptocurrency data in the database
func UpdateCryptocurrency(crypto models.Cryptocurrency) {
	// Check if the cryptocurrency already exists
	var existingCrypto models.Cryptocurrency
	err := database.DB.Where("symbol = ?", crypto.Symbol).First(&existingCrypto).Error
	if err != nil {
		// If it doesn't exist, insert it
		err = database.DB.Create(&crypto).Error
		if err != nil {
			log.Printf("Error inserting new cryptocurrency: %v", err)
		}
	} else {
		// If it exists, update it
		err = database.DB.Model(&existingCrypto).Updates(crypto).Error
		if err != nil {
			log.Printf("Error updating existing cryptocurrency: %v", err)
		}
	}
}
