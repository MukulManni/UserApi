package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getUser(c *gin.Context) {
	ID := c.Param("id")

	for _, u := range userList {
		if strconv.Itoa(u.Id) == ID {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}

	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"message": "User not found.",
		},
	)
}

func delUser(c *gin.Context) {
	ID := c.Param("id")

	for i, u := range userList {
		if strconv.Itoa(u.Id) == ID {
			userList = append(userList[:i], userList[i+1:]...)

			db.QueryRow(
				`DELETE FROM users WHERE id=$1`, u.Id,
			)

			c.IndentedJSON(
				http.StatusAccepted,
				gin.H{"message": "User deleted successfully."},
			)
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest,
		gin.H{"message": "User not found."},
	)
}

func createUser(c *gin.Context) {
	var user User

	time := time.Now()
	user.CreatedAt = time.Format("2006-01-02, 15:04:05")

	if err := c.BindJSON(&user); err == nil {
		userList = append(userList, user)

		db.QueryRow(
			`INSERT INTO users (name, dob, address, description, createdat) VALUES ($1,$2,$3,$4,$5)`,
			user.Name, user.Dob, user.Address, user.Description, user.CreatedAt,
		)

		c.IndentedJSON(http.StatusCreated, gin.H{
			"User":   user,
			"Status": "User created successfully.",
		})

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Not inserted."})
		return
	}
}

func updateUser(c *gin.Context) {
	ID := c.Param("id")

	var user User

	if err := c.BindJSON(&user); err == nil {
		for i, u := range userList {
			if strconv.Itoa(u.Id) == ID {
				if user.Id != 0 {
					userList[i].Id = user.Id
				}
				if user.Name != "" {
					userList[i].Name = user.Name
				}
				if user.Address != "" {
					userList[i].Address = user.Address
				}
				if user.Dob != "" {
					userList[i].Dob = user.Dob
				}
				if user.Description != "" {
					userList[i].Description = user.Description
				}
				if user.CreatedAt != "" {
					userList[i].CreatedAt = user.CreatedAt
				}

				db.QueryRow(
					`UPDATE users
					SET name = $1,address = $2, dob = $3, description = $4 
					WHERE id = $5;`, u.Name, u.Dob, u.Description, u.Id,
				)

				c.IndentedJSON(
					http.StatusOK,
					gin.H{
						"message": "User updated successfully.",
					},
				)

				return
			}
		}

		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "User not found",
			},
		)

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Error while updating data."})
		return
	}
}

func connectDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
    	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20),
	  	dob VARCHAR(11),
	  	address VARCHAR(50),
	  	description VARCHAR(100),
      	createdat VARCHAR(15)
    	);
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}
