package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	maxDbOpenConnection        = 5
	maxDbIdleConnection        = 2
	defaultGORMCreateBatchSize = 50
)

type DataSource struct {
	Connection *gorm.DB
}

func NewDataSource(config PostgreSQLDbConfig) *DataSource {

	dbLoggingLvl := gormlogger.Warn
	gormLogger := gormlogger.New(
		log.New(os.Stdout, "", log.LstdFlags), // io writer
		gormlogger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  dbLoggingLvl,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	dsn := "host=" + config.Url + " port=" + config.Port + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DbName
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:          gormLogger,
		CreateBatchSize: defaultGORMCreateBatchSize,
	})
	if err != nil {
		log.Panic("Failed to start db", err)
		return nil
	}

	connPool, err := db.DB()
	if err != nil {
		log.Panic("Failed to get the GORM connection pool obj", err)
		return nil
	}

	connPool.SetMaxOpenConns(maxDbOpenConnection)
	connPool.SetMaxIdleConns(maxDbIdleConnection)

	return &DataSource{
		Connection: db,
	}
}

func (ds *DataSource) Close() error {
	log.Println("Closing DB connection")
	connPool, err := ds.Connection.DB()
	if err != nil {
		log.Println("Failed to get the GORM connection pool obj")
		return err
	}
	err = connPool.Close()
	if err != nil {
		log.Println("Failed to close the DB connection")
		return err
	}
	return nil
}
