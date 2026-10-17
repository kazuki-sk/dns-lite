package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func OpenDB(driver, dsn string) (*gorm.DB, error) {
	var d gorm.Dialector

	switch driver {
	case "sqlite":
		// example: "./app.db"
		d = sqlite.Open(dsn)
	case "postgres":
		// example: "host=localhost user=app password=secret dbname=app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
		d = postgres.Open(dsn)
	case "mysql":
		// example: "app:secret@tcp(127.0.0.1:3306)/app?charset=utf8mb4&parseTime=True&loc=Local"
		d = mysql.Open(dsn)
	case "sqlserver":
		// example: "sqlserver://app:secret@localhost:1433?database=app"
		d = sqlserver.Open(dsn)
	default:
		return nil, fmt.Errorf("unknown driver: %s", driver)
	}

	return gorm.Open(d, &gorm.Config{})
}
