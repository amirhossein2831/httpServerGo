package DB

import (
	"errors"
	"github.com/amirhossein2831/httpServerGo/src/DB/msq"
	"github.com/amirhossein2831/httpServerGo/src/DB/pgs"
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/amirhossein2831/httpServerGo/src/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
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
			Logger.GetInstance().GetLogger().Error("Error in Db type ",
				zap.Error(errors.New("unsupported db type")),
				zap.Time("timestamp", time.Now()),
			)
			log.Fatal(errors.New("unsupported db type"))
		}
		db, err := dbInstance.Connect()
		if err != nil {
			Logger.GetInstance().GetLogger().Error("db connection error",
				zap.Error(errors.New("unsupported db type")),
				zap.Time("timestamp", time.Now()),
			)
			log.Fatal(err)
		}
		dbInstance.SetDb(db)
		Logger.GetInstance().GetLogger().Info("connected to db successfully",
			zap.Time("timestamp", time.Now()),
		)
	})
	return dbInstance
}
