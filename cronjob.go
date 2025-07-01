package main

import (
	User_Model "GoFr/models/user"
	"gofr.dev/pkg/gofr"
)

type User struct {
	Userid    int    `json:"id" `
	Username  string `json:"name"`
	Userphone string `json:"phone" db:"userphone"`
	Useremail string `json:"email" db:"useremail"`
}

func (u *User) TableName() string {
	return "usermanage"
}

func main() {
	app := gofr.New()

	app.GET("/hello", func(ctx *gofr.Context) (interface{}, error) {
		return "hello world", nil
	})
	app.GET("/", func(ctx *gofr.Context) (interface{}, error) {
		return "Welcome to Home Page", nil
	})
	app.AddCronJob("*/5 * * * * *", "scheduler for 1 second", func(ctx *gofr.Context) {
		ctx.Logger.Infof("scheduling the task at the time of ")

	})
	app.AddRESTHandlers(&User{})

	app.Run()
}
