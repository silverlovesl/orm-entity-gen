package models

import "fmt"

type DBConnection struct {
	Username  string
	Password  string
	Port      int
	Host      string
	Database  string
	TableName string
}

func (conn *DBConnection) GetMySQLConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&multiStatements=true&parseTime=true",
		conn.Username,
		conn.Password,
		conn.Host,
		conn.Port,
		conn.Database,
	)
}
