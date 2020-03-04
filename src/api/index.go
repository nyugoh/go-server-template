package api

import (
	. "app-template/src/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "200",
		"message": "Welcome to andromeda",
	})
}
