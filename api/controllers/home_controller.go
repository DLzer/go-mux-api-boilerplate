package controllers

import (
	"net/http"

	"github.com/DLzer/go-product-api.git/api/responses"
)

// Home response
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
