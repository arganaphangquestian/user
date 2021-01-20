package main

import (
	"log"

	"github.com/arganaphangquestian/user/repository"
	"github.com/arganaphangquestian/user/route"
)

func main() {
	repository := repository.New()
	app := route.New(repository)
	log.Fatalln(app.Listen(":8080"))
}
