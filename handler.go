package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
