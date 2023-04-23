package datasource

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/cluster05/linktree/api/config"
)

type DataSource struct {
	MySqlDB *gorm.DB
}

func Init() (*DataSource, error) {

	mySqlDB, err := setupMySQLDB(config.DatabaseConfig.MySqlDSN)
	if err != nil {
		return nil, fmt.Errorf("error opening mysqldb %w", err)
	}

	return &DataSource{
		MySqlDB: mySqlDB,
	}, nil
}
