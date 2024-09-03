package migration

import (
    "github.com/msazad/assessment/src/models"
    "github.com/msazad/assessment/utils/database"
)

func Migrate() {
    db := database.DB
    db.AutoMigrate(&models.Cryptocurrency{})
}
