package utils

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"strings"
)

// Returns a connection to db or error
// Connects to DB depending on the APP_ENV, by using APP_ENV value as a prefix to db credentials
// i.e, if
// APP_ENV=prod, it will connect to the production db, same for local, dev, staging
// APP_ENV=local, it will connect to localhost db
// APP_ENV=dev, it will connect to shared dev db
// APP_ENV=staging, it will connect to the staging db
func DbConnect() (*gorm.DB, error) {
	appEnv := strings.ToUpper(os.Getenv("APP_ENV"))

	dbDriver := os.Getenv("DB_DRIVER")
	if len(strings.TrimSpace(dbDriver)) == 0 {
		os.Setenv("DB_DRIVER", "mysql")
	}

	dbUser := os.Getenv(fmt.Sprintf("%s_DB_USER", appEnv))
	dbPass := os.Getenv(fmt.Sprintf("%s_DB_PASSWORD", appEnv))
	dbHost := os.Getenv(fmt.Sprintf("%s_DB_HOST", appEnv))
	dbPort := os.Getenv(fmt.Sprintf("%s_DB_PORT", appEnv))
	dbName := os.Getenv(fmt.Sprintf("%s_DB_NAME", appEnv))
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	if dbDriver != "mysql" {
		return nil, errors.New("only mysql driver is configured")
	}

	Log("Connecting to", strings.ToUpper(dbName), "database on", dbHost, "as", strings.ToUpper(dbUser))
	db, err := gorm.Open(dbDriver, dbUri)
	if err != nil {
		return nil, err
	}

	Log("Testing connection...")
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	Log("Connected to DB successfully...")

	Log("Setting MAX_OPEN_CONNECTIONS...")
	maxOpenCon, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTIONS"))
	if err != nil {
		LogError("Unable to read MAX_OPEN_CONNECTIONS, default value is set to 100")
		maxOpenCon = 100
	}
	db.DB().SetMaxOpenConns(maxOpenCon)

	Log("Setting MAX_IDLE_CONNECTIONS...")
	maxIdleCon, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTIONS"))
	if err != nil {
		LogError("Unable to read MAX_IDLE_CONNECTIONS, default value is set to 64")
		maxIdleCon = 64
	}
	db.DB().SetMaxIdleConns(maxIdleCon)

	return db, nil
}

func RedisConnect() (client *redis.Client) {
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
	Log("Connecting to redis...")
	Log("Connected to redis...")
	return client
}

func AutoMigrateDB(db *gorm.DB) {
	Log("Running auto migrations...")
	db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate()
	Log("Done migrating...")
}

func SeedDB(db *gorm.DB) {}
