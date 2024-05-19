package config

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() (*gorm.DB, error) {
	dbType := os.Getenv("DB_TYPE")

	if dbType == "postgres" {
		return postgresConnection()
	} else if dbType == "mysql" {
		return mysqlConnection()
	}
	return nil, errors.New("the db type is not supported")
}

func mysqlConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	port := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func postgresConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	port := os.Getenv("PG_PORT")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", host, user, pass, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
