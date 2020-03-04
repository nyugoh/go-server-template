package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

func InitLogger() {
	logFolder := os.Getenv("LOG_FOLDER")
	env := os.Getenv("ENV")
	appName := os.Getenv("APP_NAME")
	if env == "dev" {
		pwd, err := os.Getwd()
		if err == nil {
			logFolder = fmt.Sprintf("%s/logs/", pwd)
		}
	}
	if len(logFolder) == 0 {
		logFolder = "logs"
	}
	if len(appName) == 0 {
		appName = "app-logs"
	}
		writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s.log", logFolder+appName+"-old", "%Y-%m-%d"),
		rotatelogs.WithLinkName(logFolder+appName+".log"),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		fmt.Println("Failed to initialize log file ", err.Error())
	}
	log.SetOutput(writer)
	return
}

func Log(msg ...interface{}) {
	fmt.Printf("%s: ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(msg...)
	log.Println(msg...)
}
