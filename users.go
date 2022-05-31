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

func fetchUsers() ([]User, error) {
	rows, err := db.Query(
		`SELECT * FROM users`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]User, 0, 10)

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
