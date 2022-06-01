package main

func initializeRoutes() {

	r.GET("/", mainPage)

	r.GET("/get/:id", getUser)
	r.DELETE("/delete/:id", delUser)

	r.POST("/create", createUser)
	r.PUT("/update/:id", updateUser)
}
