package logger

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"github.com/cluster05/linktree/api/config"
)

func DBLogger() logger.Interface {

	logLevel := logger.Silent
	if config.AppConfig.Env == "dev" {
		logLevel = logger.Info
	}

	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
}
