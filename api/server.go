package api

import (
	"fmt"
	"log"
	"os"

	"github.com/IhsanBhee/mu-golang/api/controllers"
	"github.com/IhsanBhee/mu-golang/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))

	seed.Load(server.DB)

	server.Run(":8080")
}
