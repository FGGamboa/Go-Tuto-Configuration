package database

import (
	"context"
	"database/sql"
	"fmt"
	"tuto-configuration/configuration"

	_ "github.com/go-sql-driver/mysql"
)

func CreateConnection(ctx context.Context, config *configuration.Configuration) *sql.DB {

	connectionString := fmt.Sprintf(
		"%s:123456@tcp(%s:%d)/%s",
		config.DB.User,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db
}
