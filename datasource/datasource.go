package datasource

import (
	"fmt"

	"github.com/cluster05/linktree/api/config"
	"gorm.io/gorm"
)

type DataSource struct {
	MySqlDB *gorm.DB
}

func Init() (*DataSource, error) {

	mySqlDB, err := setupMySqlDB(config.DatabaseConfig.MySqlDNS)
	if err != nil {
		return nil, fmt.Errorf("error opening mysqldb %w", err)
	}

	return &DataSource{
		MySqlDB: mySqlDB,
	}, nil
}
