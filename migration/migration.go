package migration

import (
    "web-api/src/models"
    "web-api/utils/database"
)

func Migrate() {
    db := database.DB
    db.AutoMigrate(&models.Cryptocurrency{})
}
