package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/silverlovesl/orm-entity-gen/models"
)

type TableService struct {
	db *sql.DB
}

func NewTableService(db *sql.DB) *TableService {
	return &TableService{
		db: db,
	}
}

// ListAllTables list all table name
func (serv *TableService) ListAllTables() ([]string, error) {
	result := []string{}
	rows, err := serv.db.Query("show tables")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tableName := ""
		rows.Scan(&tableName)
		result = append(result, tableName)
	}
	return result, nil
}

// GetColumnDescription
func (serv *TableService) GetColumnDescription(tableName string) ([]models.ColumnDesc, error) {
	rows, err := serv.db.Query(fmt.Sprintf("desc %s", tableName))
	if err != nil {
		log.Fatalf("Error %s", err.Error())
		return nil, err
	}
	result := []models.ColumnDesc{}
	for rows.Next() {
		col := models.ColumnDesc{}
		if err := rows.Scan(
			&col.Field,
			&col.Type,
			&col.Null,
			&col.Key,
			&col.Default,
			&col.Extra,
		); err != nil {
			log.Fatalf("Read column def %s", err.Error())
		}
		result = append(result, col)
	}
	return result, nil
}
