package main

import (
	"fmt"
	"german-holidays/entities"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	DB_HOST := os.Getenv("POSTGRES_CONTAINER_NAME")
	DB_PORT := os.Getenv("POSTGRES_INTERNAL_PORT")
	DB_NAME := os.Getenv("POSTGRES_DB")
	DB_USER := os.Getenv("POSTGRES_USER")
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: dbLogger})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfully connected to the database!")
	err = db.AutoMigrate(&entities.Region{}, &entities.Holiday{})

	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "App is up and running!",
		})
	})

	INTERFACE := ":"
	PORT := os.Getenv("GO_INTERNAL_PORT")
	ADDRESS := INTERFACE + PORT
	err = router.Run(ADDRESS)
	if err != nil {
		log.Fatal("Failed to start server!")
	}

}
