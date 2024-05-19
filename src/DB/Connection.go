package DB

import (
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/DB/msq"
	"github.com/amirhossein2831/httpServerGo/src/DB/pgs"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	dbInstance Database
	once       sync.Once
)

type Database interface {
	Connect() (*gorm.DB, error)
	GetDb() *gorm.DB
	SetDb(db *gorm.DB)
}

func GetInstance() Database {
	once.Do(func() {
		dbType := config.GetInstance().Get("DB_TYPE")
		switch dbType {
		case "postgres":
			dbInstance = &pgs.PostgresDB{}
		case "mysql":
			dbInstance = &msq.MysqlDB{}
		default:
			log.Fatal(errors.New("unsupported db type"))
		}
		db, err := dbInstance.Connect()
		if err != nil {
			log.Fatal(err)
		}
		dbInstance.SetDb(db)
		println("connected to DB successfully")
	})
	return dbInstance
}
