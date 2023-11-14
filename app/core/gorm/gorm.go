package gorm

import (
	"AILN/app/common"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

// Assuming common and newLogger functions are defined elsewhere in your code

func NewGorm() (*gorm.DB, error) {
	var (
		logout io.Writer
		err    error
	)

	logout = os.Stdout
	if common.CONFIG.String("gorm_log.outType") == "file" {
		logout, err = os.Open(common.CONFIG.String("gorm_log.out"))
		if err != nil {
			return nil, fmt.Errorf("NewGorm: Cannot open file %s: %v", common.CONFIG.String("gorm_log.out"), err)
		}
	}

	newLogger := logger.New(
		log.New(logout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	config := common.CONFIG.StringMap("mysql")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", config["user"], config["password"], config["server"], config["port"], config["database"], config["config"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(common.CONFIG.Int("mysql.maxOpenConns"))
	sqlDB.SetMaxIdleConns(common.CONFIG.Int("mysql.maxIdleConns"))
	return db, nil
}
