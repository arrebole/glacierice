package service

import (
	"github.com/jinzhu/gorm"
	// postgres 驱动
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=111.231.86.218 port=5432 user=siteowner dbname=site_assets password=goforit")
	if err != nil {
		panic(err.Error())
	}
	return db
}
