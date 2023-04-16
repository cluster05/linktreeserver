package datasource

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/cluster05/linktree/util/logger"
)

var (
	sqlConnection      *gorm.DB
	sqlConnectionError error
	sqlOnce            sync.Once
)

func setupMySqlDB(sqlDNS string) (*gorm.DB, error) {

	sqlOnce.Do(func() {
		connection, err := gorm.Open(mysql.Open(sqlDNS), &gorm.Config{Logger: logger.DBLogger()})
		if err != nil {
			sqlConnectionError = err
		} else {
			sqlConnection = connection
		}

	})
	return sqlConnection, sqlConnectionError
}
