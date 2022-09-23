package store

import "database/sql"

type MysqlStorage struct {
	DB *sql.DB
}

func NewMySQLStorage(db *sql.DB) MysqlStorage{
	return  MysqlStorage{
		DB: db,
	}
}