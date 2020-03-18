package controllers

import (
	"net/http"

	"github.com/IhsanBhee/mu-golang/api/responses"
)

func (server *Server) home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
