package api

import (
	. "app-template/app/utils"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

type App struct {
	DB    *gorm.DB
	Name  string
	Port  string
	Redis redis.Client
}

func (app *App) LoadApiInitialData() {
	Log("Loading API initial data")
}
