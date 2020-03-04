package main

import (
	"app-template/src/api"
	"app-template/src/utils"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Unable to read .env file", err.Error())
	}
	port := ":5000"
	if len(strings.TrimSpace(os.Getenv("PORT"))) > 0 {
		port = ":" + os.Getenv("PORT")
	}

	if len(strings.TrimSpace(os.Getenv("APP_NAME"))) == 0 {
		os.Setenv("APP_NAME", "GO_APP_API")
	}

	utils.InitLogger()
	utils.Log("Starting app...")

	r := mux.NewRouter()
	r.HandleFunc("/", api.Home).Methods("GET")


	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	utils.Log(fmt.Sprintf("Magic brewing on port %s", port))
	if err := http.ListenAndServe(port, loggedRouter); err != nil {
		utils.Log("App terminated: ", err.Error())
	}
}