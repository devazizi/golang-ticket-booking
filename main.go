package main

import (
	"fmt"
	"log"
	"os"

	"github.com/devazizi/golang-ticket-booking.git/controllers"
	"github.com/devazizi/golang-ticket-booking.git/services"
	"github.com/labstack/echo/v4"
)

func router(e echo.Echo) {
	e.GET("/", controllers.IndexTickets)
}

func setDotEnvFileVars() {
	appDirectory, err := os.Getwd()

	if err != nil {
		log.Fatalln("fail to fetch directory")
	}

	var envFilePath string = appDirectory + "/.env"
	if _, err := os.Stat(envFilePath); err == nil {
		services.SetEnvVariables(envFilePath)
	}
}

func main() {

	setDotEnvFileVars()

	fmt.Println(os.Getenv("Foo"))
	echo := echo.New()

	router(*echo)

	echo.Logger.Fatal(echo.Start(":8080"))
}
