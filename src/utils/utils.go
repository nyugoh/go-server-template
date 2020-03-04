package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]interface{}{"error": msg, "status": 0})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	Log(fmt.Sprintf("RESPONSE:: Status:%d Payload: %v", code, payload))
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
