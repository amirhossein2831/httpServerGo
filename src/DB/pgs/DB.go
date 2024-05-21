package pgs

import (
	"fmt"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func (pg *PostgresDB) Connect() (*gorm.DB, error) {
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

func (pg *PostgresDB) GetDb() *gorm.DB {
	return pg.db
}

func (pg *PostgresDB) SetDb(db *gorm.DB) {
	pg.db = db
}
