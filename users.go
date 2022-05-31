package main

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

var userList []User
