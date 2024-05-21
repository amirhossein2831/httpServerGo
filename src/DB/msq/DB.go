package msq

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	db *gorm.DB
}

func (ms *MysqlDB) Connect() (*gorm.DB, error) {
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

func (ms *MysqlDB) GetDb() *gorm.DB {
	return ms.db
}

func (ms *MysqlDB) SetDb(db *gorm.DB) {
	ms.db = db
}
