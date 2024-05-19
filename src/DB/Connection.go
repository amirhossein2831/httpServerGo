package DB

import (
	"errors"
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbType := config.GetInstance().Get("DB_TYPE")

	if dbType == "postgres" {
		return postgresConnection()
	} else if dbType == "mysql" {
		return mysqlConnection()
	}

	return nil, errors.New("the db type is not supported")
}

func mysqlConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetInstance().Get("MYSQL_USER"),
		config.GetInstance().Get("MYSQL_PASS"),
		config.GetInstance().Get("DB_HOST"),
		config.GetInstance().Get("MYSQL_PORT"),
		config.GetInstance().Get("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func postgresConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		config.GetInstance().Get("DB_HOST"),
		config.GetInstance().Get("PG_USER"),
		config.GetInstance().Get("PG_PASS"),
		config.GetInstance().Get("DB_NAME"),
		config.GetInstance().Get("PG_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
