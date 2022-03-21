package controllers

import (
	"net/http"

	"github.com/hugoalmeidahh/go-restful/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this Awesome API created in GoLang!")

}
