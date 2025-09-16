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
	"strings"
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
	dbName := config["database"]

	// 1. Connect to MySQL server without specifying a database
	dsnWithoutDB := fmt.Sprintf("%v:%v@tcp(%v:%v)/?%v", config["user"], config["password"], config["server"], config["port"], config["config"])

	tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL server to create database: %v", err)
	}

	// Ensure the temporary connection is closed
	sqlTempDB, err := tempDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying SQL DB from temporary connection: %v", err)
	}
	defer sqlTempDB.Close()

	// 2. Create the database if it doesn't exist
	createDBStmt := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)
	if err := tempDB.Exec(createDBStmt).Error; err != nil {
		return nil, fmt.Errorf("failed to create database %s: %v", dbName, err)
	}

	// 3. Now, connect to the specific database
	dsnWithDB := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", config["user"], config["password"], config["server"], config["port"], dbName, config["config"])

	db, err := gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// 4. Execute lab.sql to initialize tables and data
	if err := executeSQLFile(db, "deploy/db/lab.sql"); err != nil {
		return nil, fmt.Errorf("failed to execute lab.sql: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(common.CONFIG.Int("mysql.maxOpenConns"))
	sqlDB.SetMaxIdleConns(common.CONFIG.Int("mysql.maxIdleConns"))
	return db, nil
}

func executeSQLFile(db *gorm.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	statements := strings.Split(string(content), ";")
	for _, stmt := range statements {
		trimmedStmt := strings.TrimSpace(stmt)
		if trimmedStmt == "" {
			continue
		}
		// Skip CREATE DATABASE and USE statements as they are handled separately
		if strings.HasPrefix(strings.ToUpper(trimmedStmt), "CREATE DATABASE") || strings.HasPrefix(strings.ToUpper(trimmedStmt), "USE") {
			continue
		}
		if err := db.Exec(trimmedStmt).Error; err != nil {
			return fmt.Errorf("error executing SQL statement: %s, error: %v", trimmedStmt, err)
		}
	}
	return nil
}
