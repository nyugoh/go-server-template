package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func InitLogger() error {
	// Set log folder to be the current project root if app env is DEV/LOCAL or LOG_FOLDER is not set
	logFolder := os.Getenv("LOG_FOLDER")
	appEnv := strings.TrimSpace(strings.ToUpper(os.Getenv("APP_ENV")))
	if appEnv == "DEV" || appEnv == "LOCAL" || len(logFolder) == 0 {
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}
		logFolder = fmt.Sprintf("%s/logs/", pwd)
	}

	// Set log file names according to app_name
	appName := strings.TrimSpace(strings.ToLower(os.Getenv("APP_NAME")))
	if len(appName) == 0 {
		appName = "app-logs" // Default log file name
	}
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s.log", logFolder+appName+"-old", "%Y-%m-%d"),
		rotatelogs.WithLinkName(logFolder+appName+".log"),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithRotationCount(10000),
	)
	if err != nil {
		LogError("Failed to initialize log file ", err.Error())
		return err
	}
	log.SetOutput(writer)
	return nil
}

func Log(msg ...interface{}) {
	LogInfo(msg)
}

func LogInfo(msg ...interface{}) {
	msgStr := removeBraces(msg)
	fmt.Printf(fmt.Sprintf("%s: INFO: %v\n", CurrentTime(), msgStr))
	log.Println("INFO: ", msgStr)
}

func LogWarning(msg ...interface{}) {
	msgStr := removeBraces(msg)
	fmt.Printf(fmt.Sprintf("%s: WARNING: %v\n", CurrentTime(), msgStr))
	log.Println("WARNING: ", msgStr)
}

func LogError(msg ...interface{}) {
	msgStr := removeBraces(msg)
	fmt.Printf(fmt.Sprintf("%s: ERROR: %v\n", CurrentTime(), msgStr))
	log.Printf(fmt.Sprintf("ERROR: %v\n", msgStr))
}

func removeBraces(msg []interface{}) string {
	msgStr := fmt.Sprintf("%v", msg)
	msgStr = strings.Replace(msgStr, "[", "", 2)
	msgStr = strings.Replace(msgStr, "]", "", 2)
	return msgStr
}
