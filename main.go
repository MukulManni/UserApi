package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var r *gin.Engine
var db *sql.DB

func main() {
	url, ok := os.LookupEnv("DATABASE_URL")

	if !ok {
		log.Fatalln("Database url is required.")
	}

	var err error
	db, err = connectDB(url)

	if err != nil {
		log.Fatalf("Error connecting database: %s", err.Error())
	}

	port := os.Getenv("PORT")
	//port = "80"

	users, err := fetchUsers()

	if err != nil {
		log.Fatalln("Unable to retrieve messages from database.")
	}

	userList = append(userList, users...)

	r = gin.Default()

	initializeRoutes()

	r.Run(":" + port)
}
