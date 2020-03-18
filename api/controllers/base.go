package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/IhsanBhee/mu-golang/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(DbDriver, DbHost, DbPort, DbName, DbUser, DbPass string) {
	var err error

	if DbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s dbname=%s user=%s pass=%s sslmode=disable",
			DbHost, DbPort, DbName, DbUser, DbPass)
		server.DB, err = gorm.Open(DbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listener to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
