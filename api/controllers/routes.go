package controllers

import (
	"github.com/IhsanBhee/mu-golang/api/middlewares"
)

func (s *Server) initializeRoutes() {
	//Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.home)).Methods("GET")

	//Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.login)).Methods("POST")

	//User Route
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.save)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.findAll)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.findById)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareJSON(s.update))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.delete)).Methods("DELETE")
}
