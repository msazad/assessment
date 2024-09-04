package main

import (
    "fmt"
    "log"
    "github.com/msazad/assessment/migration"
    "github.com/msazad/assessment/route"
    "github.com/msazad/assessment/src/cronjob"
    "github.com/msazad/assessment/utils/database"
    "github.com/msazad/assessment/utils/middleware"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func initConfig() {
    // Load configuration from .env file
    viper.SetConfigFile(".env")
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
}

func main() {
    // Initialize configuration
    initConfig()

    // Start cron job
    cronjob.StartCronJob()

    // Initialize PostgreSQL database connection
    database.Initpostgres()

    // Run migrations
    migration.Migrate()

    // Set up Gin router
    router := gin.Default()

    // Set up middleware
    router.Use(middleware.TracingMiddleware())

    // Set up CORS
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowAllOrigins = true
    router.Use(cors.New(corsConfig))

    // Set up routes
    route.SetupRoutes(router)

    // Start server
    port := viper.GetString("SERVER_PORT")
    if port == "" {
        port = ":8080" // default port if not specified
    }
    fmt.Printf("Server started at port %s\n", port)
    if err := router.Run(port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

