package main

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

var userList []User //The local list for users

//fetchUsers() => fetch user data from database and put in local user list

func fetchUsers() ([]User, error) {

	//sql query to get data of users
	rows, err := db.Query(
		`SELECT * FROM users`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close() //Closing current row at the end of fucntion

	users := make([]User, 0, 10)

	//Scanning and adding each row's data in a local user slice and then return that slice as the updated user data
	for rows.Next() {
		u := User{}

		err = rows.Scan(&u.Id, &u.Name, &u.Dob, &u.Address, &u.Description, &u.CreatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil

}
