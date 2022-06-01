package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// getUser() => Handler for the get request for getting user data

func getUser(c *gin.Context) {
	ID := c.Param("id") //store the parameter value from url in a local variable

	//search in the the updated local userlist for the id
	//if matched return ok status and json user data
	for _, u := range userList {
		if strconv.Itoa(u.Id) == ID {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}

	//If no user found return appropriate status and message
	c.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"message": "User not found.",
		},
	)
}

// delUser() => Handler for the delete request for deleting a user

func delUser(c *gin.Context) {
	ID := c.Param("id") //store the parameter value from url in a local variable

	//Search id in the local user list
	for i, u := range userList {
		if strconv.Itoa(u.Id) == ID {
			userList = append(userList[:i], userList[i+1:]...) //If matched removed that user from local list

			//Perform sql query to remove user from database
			db.QueryRow(
				`DELETE FROM users WHERE id=$1`, u.Id,
			)

			//If user found and data updated, return appropriate status and message
			c.IndentedJSON(
				http.StatusAccepted,
				gin.H{"message": "User deleted successfully."},
			)
			return
		}
	}

	//If no user found return appropriate status and message
	c.IndentedJSON(http.StatusBadRequest,
		gin.H{"message": "User not found."},
	)
}

// creteUser() => Handler for the post request to create new user

func createUser(c *gin.Context) {
	var user User //A local user copy for storing temporary user data

	time := time.Now()                                   //Get local time from the server
	user.CreatedAt = time.Format("2006-01-02, 15:04:05") //Format time in date and time and store in the local user copy

	//Bind the posted user data through request to the local user struct
	if err := c.BindJSON(&user); err == nil {

		//If no error found, perform sql query to add posted data to database
		row := db.QueryRow(
			`INSERT INTO users (name,dob,address,description,createdat) VALUES ($1,$2,$3,$4,$5) RETURNING id;`,
			user.Name, user.Dob, user.Address, user.Description, user.CreatedAt,
		)

		err := row.Scan(&user.Id) //get the created id from database

		if err != nil {
			fmt.Println(err)
		}

		userList = append(userList, user) //Append the new formed user in the local user list

		//Return appropriate status and data in json form
		c.IndentedJSON(http.StatusCreated, gin.H{
			"User":   user,
			"Status": "User created successfully.",
		})

	} else {
		//In case of error return badReques and not inserted message
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Not inserted."})
		return
	}
}

// updateUser() => Handler for the put request to update user data

func updateUser(c *gin.Context) {
	ID := c.Param("id") //store the parameter value from url in a local variable

	var user User //A local user copy for storing temporary user data

	//Bind the posted user data through request to the local user struct
	if err := c.BindJSON(&user); err == nil {

		//Search for the requested user in userlist
		for i, u := range userList {

			//If found, then update the changed fields in the local user list
			if strconv.Itoa(u.Id) == ID {
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

				//Perform sql query to update data of requested user in database
				db.QueryRow(
					`UPDATE users SET name=$1,address=$2,dob=$3,description=$4 WHERE id = $5 RETURNING id;`,
					userList[i].Name, userList[i].Address, userList[i].Dob, userList[i].Description, userList[i].Id,
				)

				//Return StatusOk and Updated successfully message
				c.IndentedJSON(
					http.StatusOK,
					gin.H{
						"message": "User updated successfully.",
					},
				)

				return
			}
		}

		//If no user found return badRequest and user not found message
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "User not found",
			},
		)

	} else {

		//If unable to bind data then, return appropritate message
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Error while updating data."})
		return
	}
}

// connectDB() => Function to connect App with the database

func connectDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL) //Open database from the database url fetched from heroku

	if err != nil {
		return nil, err
	}

	err = db.Ping() //Check the connection with database

	if err != nil {
		return nil, err
	}

	//If connected, then execute sql query to create new table
	_, err = db.Exec(`
    	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(20),
	  	dob VARCHAR(11),
	  	address text,
	  	description text,
      	createdat text
    	);
  `)

	if err != nil {
		return nil, err
	}

	return db, nil //return the database as a variable
}
