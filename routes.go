package main

func initializeRoutes() {
	r.GET("/get/:id", getUser)
	r.DELETE("/delete/:id", delUser)

	r.POST("/create", createUser)
	r.PUT("/update/:id", updateUser)
}
