package datasource

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	sqlConnection      *sqlx.DB
	sqlConnectionError error
	sqlOnce            sync.Once
)

func setupMySqlDB(sqlDNS string) (*sqlx.DB, error) {

	fmt.Println("Connecting ", sqlDNS)

	sqlOnce.Do(func() {
		connection, err := sqlx.Open("mysql", sqlDNS)
		if err != nil {
			sqlConnectionError = err
		} else {

			connection.SetMaxIdleConns(100)
			connection.SetMaxOpenConns(1000)
			connection.SetConnMaxLifetime(4 * time.Hour)

			sqlConnection = connection

			//setupDBSchema()
		}

	})
	return sqlConnection, sqlConnectionError
}

func setupDBSchema() {

}
