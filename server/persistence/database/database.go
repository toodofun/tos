package database

import (
	"encoding/json"
	"errors"
	"github.com/MR5356/tos/config"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var (
	once sync.Once
	db   *Database
)

type Database struct {
	*gorm.DB
}

var (
	DBDriverNotSupport = errors.New("database driver not support")
)

func NewDatabase(cfg *config.Config) *Database {
	once.Do(func() {
		var err error
		db, err = initDB(cfg)
		if err != nil {
			logrus.Fatalf("Failed to initialize database: %v", err)
		}
	})
	return db
}

func GetDB() *Database {
	//once.Do(func() {
	//	err := retry.Do(
	//		func() (err error) {
	//			db, err = initDB(config.Current())
	//			return err
	//		},
	//		retry.Attempts(config.DefaultRetryCount),
	//		retry.Delay(config.DefaultRetryDelay),
	//		retry.LastErrorOnly(true),
	//		retry.DelayType(retry.DefaultDelayType),
	//		retry.OnRetry(func(n uint, err error) {
	//			logrus.Warnf("[%d/%d]: retry to initialize database: %v", n+1, config.DefaultRetryCount, err)
	//		}),
	//	)
	//	if err != nil {
	//		logrus.Fatalf("Failed to initialize database: %v", err)
	//	}
	//})
	return db
}

func initDB(cfg *config.Config) (database *Database, err error) {
	var driver gorm.Dialector
	logrus.Infof("database driver: %s, dsn: %s", cfg.Persistence.Database.Driver, cfg.Persistence.Database.DSN)
	switch cfg.Persistence.Database.Driver {
	case "sqlite":
		driver = sqlite.Open(cfg.Persistence.Database.DSN)
	default:
		return nil, DBDriverNotSupport
	}

	var dbLogLevel = logger.Error
	if cfg.Server.Debug {
		dbLogLevel = logger.Info
	}
	logrus.Debugf("database log level: %+v", dbLogLevel)

	client, err := gorm.Open(driver, &gorm.Config{
		Logger: logger.Default.LogMode(dbLogLevel),
	})
	if err != nil {
		return nil, err
	}

	db, err := client.DB()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.Persistence.Database.MaxIdleConn)
	db.SetMaxOpenConns(cfg.Persistence.Database.MaxOpenConn)
	db.SetConnMaxLifetime(cfg.Persistence.Database.ConnMaxLift)
	db.SetConnMaxIdleTime(cfg.Persistence.Database.ConnMaxIdle)

	dbStat, _ := json.Marshal(db.Stats())
	logrus.Infof("database stats: %s", dbStat)
	return &Database{client}, nil
}
