package utils

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
)

// Returns a connection to db or error
func DbConnect() (*sql.DB, error) {
	dbUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver != "mysql" {
		Log("Only mysql driver is configured.")
		log.Fatal("Exiting app...")
	}
	db, err := sql.Open(dbDriver, dbUri)
	if err != nil {
		Log("unable to connect to DB", err.Error())
		return nil, err
	} else {
		Log("Connected to DB")
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(64)
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