package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var DB *gorm.DB

func initpostgres() {
	once.Do(func() {
		user := viper.GetString("DB_USER")
		password := viper.GetString("DB_PASSWORD")
		host := viper.GetString("DB_HOST")
		port := viper.GetString("DB_PORT")
		dbname := viper.GetString("DB_NAME")

		//constuct the data source name(DSN)
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host,user,password,dbname,port)

			//Establish database connection
			db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
			if err !=nil{
				log.Fatalf("Failed to connect to database :%v",err)
			}
			//Assign the conncected database instance to the global variable
			DB=db
	})
}
