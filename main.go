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
	url, ok := os.LookupEnv("DATABASE_URL") //Database url for postgres database of heroku

	if !ok {
		log.Fatalln("Database url is required.")
	}

	var err error
	db, err = connectDB(url) //function defined in handler.go

	if err != nil {
		log.Fatalf("Error connecting database: %s", err.Error())
	}

	port := os.Getenv("PORT") //to get port number from heroku (heroku dependency)
	//port = "80"

	users, err := fetchUsers() //function defined in users.go

	if err != nil {
		log.Fatalln("Unable to retrieve messages from database.")
	}

	userList = append(userList, users...) //update local userlist with database

	r = gin.Default() //intialize default gin router

	initializeRoutes() //function defined in routes.go

	r.Run(":" + port)
}
