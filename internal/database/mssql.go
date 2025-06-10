package database

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"querybuilder/internal/config"
)

func NewMssqlStorage(cnf config.DBConfig) (*sqlx.DB, error) {
	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;encrypt=disable",
		cnf.Server, cnf.Port, cnf.User, cnf.Pass, cnf.Database)
	db, err := sqlx.Open("mssql", connString)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %v", err.Error())
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err.Error())
	}
	fmt.Println("Successfully connected to SQL Server 2005!")
	return db, nil
}
