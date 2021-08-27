package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

//ConnectGorm Database Connection to Gorm V2
func ConnectGorm() {
	databaseConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	var err error
	db, err = gorm.Open(mysql.Open(databaseConfig), initConfig())

	if err != nil {
		fmt.Println(err)
		panic("Fail To Connect Database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(20)
}

//initConfig Initialize Config
func initConfig() *gorm.Config {
	return &gorm.Config{
		Logger:         initLog(),
		NamingStrategy: initNamingStrategy(),
	}
}

//initLog Connection Log Configuration
func initLog() logger.Interface {
	f, _ := os.Create("gorm.log")
	newLogger := logger.New(log.New(io.MultiWriter(f, os.Stdout), "\r\n", log.LstdFlags), logger.Config{
		Colorful:      true,
		LogLevel:      logger.Info,
		SlowThreshold: time.Second,
	})
	return newLogger
}

//initNamingStrategy Init NamingStrategy
func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: true,
		TablePrefix:   "",
	}
}
