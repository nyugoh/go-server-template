package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func SendError(c *gin.Context, msg string) {
	LogError(msg)
	c.JSON(http.StatusBadRequest, gin.H{
		"error": msg,
	})
}

func SendJson(c *gin.Context, payload gin.H) {
	c.JSON(http.StatusOK, payload)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]interface{}{"error": msg})
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

func ValidateEmail(email string) (bool, error) {
	var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) > 254 || !rxEmail.MatchString(email) {
		return false, errors.New("email is invalid")
	}
	return true, nil
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		Log("Hit an auth required endpoint...")
		session := sessions.Default(c)
		userId := session.Get("userId")
		username := session.Get("username")
		if username == nil || userId == nil {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Redirect(302, "/auth/login")
			return
		}
		Log("Request made by user Id:", userId, "Username:", username)
		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}

func MetricsMonitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		// access the status we are sending
		status := c.Writer.Status()
		log.Println("Request: ", c.Request.URL.Path, "took:", latency, "Status:", status)
	}
}

func ExitApp(code int) {
	Log("Exiting app...")
	os.Exit(code)
}
