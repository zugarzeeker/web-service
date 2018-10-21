package main

import (
	"fmt"
	"log"
	"os"

	"github.com/expenseledger/web-service/controller"
	"github.com/expenseledger/web-service/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	var dbinfo string
	if os.Getenv("MODE") == "DEVELOPMENT" {
		dbinfo = fmt.Sprintf(
			"user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
	} else {
		dbinfo = os.Getenv("DATABASE_URL")
	}

	err = database.Init(dbinfo)
	if err != nil {
		log.Fatal("Error initializing database")
	}

	err = database.CreateTables()
	if err != nil {
		log.Fatal("Error creating tables")
	} else {
		log.Println("Successfully created tables")
	}

	router := controller.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Error running the server", err)
	}
}
