package app

import (
	"app-template/app/api"
	. "app-template/app/utils"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	r   *gin.Engine
	app api.App
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to read .env file", err.Error())
	}

	// Set up logger
	if err := InitLogger(); err != nil {
		LogError(err.Error())
		ExitApp(1)
	}

	Log("Initializing app...")

	// Set app port
	app.Port = ":5000"                                 // Default value
	if len(strings.TrimSpace(os.Getenv("PORT"))) > 0 { // Check .env file
		Log("Setting app port...")
		app.Port = ":" + os.Getenv("PORT")

	}
	Log("Port set to", app.Port)

	// Set app name
	if len(strings.TrimSpace(os.Getenv("APP_NAME"))) == 0 {
		Log("No app name set in .env file...")
		Log("Setting default app name...")
		os.Setenv("APP_NAME", "qpos")
	}
	app.Name = os.Getenv("APP_NAME")
	Log("App name is:", app.Name)

	Log("Done initializing app...")

	appEnv := strings.ToLower(os.Getenv("APP_ENV"))
	if len(strings.TrimSpace(appEnv)) == 0 {
		Log("App env has not been set...")
		os.Setenv("APP_ENV", "DEV")
		Log("Setting to development...")
	}
	Log("App is running in", strings.ToUpper(appEnv), "mode...")

	// Setting app mode to gin.ReleaseMode unless in DEV or LOCAL env
	if appEnv != "dev" && appEnv != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	Log("Gin server is running in", gin.Mode(), "mode")

	// Connect to DB
	db, err := DbConnect()
	if err != nil {
		LogError("unable to connect to DB:", err.Error())
		ExitApp(1)
	}
	app.DB = db

	// Init DB, update table structures
	AutoMigrateDB(app.DB)

	if ok, err := strconv.ParseBool(os.Getenv("INIT_DB")); ok && err == nil {
		SeedDB(app.DB) // Insert initial data required by the app to start
	}

	app.LoadApiInitialData()

}

func Run() {
	// Make a gin router
	r = gin.Default()

	// Session to use in auth
	r.Use(sessions.Sessions("qpos_session", sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))

	// Server static files
	r.Static("/assets", "./assets")

	// Load all templates
	r.LoadHTMLGlob("templates/*")

	// Register all routes
	initRoutes()

	// Close DB on app close or panic
	defer app.DB.Close()

	Log("Starting app...")
	Log(fmt.Sprintf("Magic brewing on port %s", app.Port))
	if err := r.Run(app.Port); err != nil {
		LogError("App terminated: ", err.Error())
	}
}

func initRoutes() {
	// Add middleware to monitor all request, /metric endpoint for analytics
	r.Use(MetricsMonitor())

	r.GET("/", app.HomePage)
	r.GET("/api/v1", app.IndexPage)

	// Add your other routes here

}
