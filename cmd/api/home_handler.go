package main

import (
	"fmt"
	"net/http"

	"github.com/italolelis/hellowork/config"
	"github.com/italolelis/hellowork/response"
)

func Home(app config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, fmt.Sprintf("Welcome to %s, this is version %s", app.Name, app.Version))
	}
}
