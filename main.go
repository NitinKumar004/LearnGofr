package main

import (
	"gofr.dev/pkg/gofr"

	"GoFr/database"
	taskHandler "GoFr/handler/task"
	userHandler "GoFr/handler/user"

	taskService "GoFr/service/task"
	userService "GoFr/service/user"

	taskStore "GoFr/store/task"
	userStore "GoFr/store/user"
)

func main() {
	// Create Gofr app
	app := gofr.New()

	// Connect DB
	db, err := database.Databasconnection()
	if err != nil {

		return
	}

	// task Layer Setup
	tStore := taskStore.New(db)
	tService := taskService.New(tStore)
	tHandler := taskHandler.New(tService)

	//user Layer Setup
	uStore := userStore.New(db)
	uService := userService.New(uStore)
	uHandler := userHandler.New(uService)

	// task Routes
	app.POST("/task", tHandler.Addtask)
	app.GET("/task", tHandler.GetAllTask)
	app.GET("/task/{id}", tHandler.GetTaskById)
	app.PATCH("/task/{id}", tHandler.CompleteTask)
	app.DELETE("/task/{id}", tHandler.DeleteTask)
	//
	//// user Routes
	app.POST("/user", uHandler.AddUser)
	app.GET("/user", uHandler.GetAllUsers)
	app.GET("/user/{id}", uHandler.GetUserByID)
	app.DELETE("/user/{id}", uHandler.DeleteUserByID)
	app.DELETE("/user", uHandler.DeleteAllUsers)

	// Start the server
	app.Run()
}
