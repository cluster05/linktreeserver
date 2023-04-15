package datasource

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/cluster05/linktree/api/config"
)

type DataSource struct {
	MySqlDB *sqlx.DB
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
