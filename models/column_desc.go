package models

import (
	"database/sql"
	"strings"
)

type ColumnDesc struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default sql.NullString
	Extra   string
}

func (col *ColumnDesc) IsPK() bool {
	return col.Key == "PRI"
}

func (col *ColumnDesc) IsNullable() bool {
	return col.Null == "YES"
}

func (col *ColumnDesc) Get2GoType() string {
	if strings.HasPrefix(col.Type, "varchar") ||
		strings.HasPrefix(col.Type, "char") ||
		strings.HasPrefix(col.Type, "text") ||
		strings.HasPrefix(col.Type, "mediumtext") ||
		strings.HasPrefix(col.Type, "longtext") {
		if col.IsNullable() {
			return "null.String"
		}
		return "string"
	}

	if strings.HasPrefix(col.Type, "int") ||
		strings.HasPrefix(col.Type, "smallint") ||
		strings.HasPrefix(col.Type, "bigint") {
		if col.IsNullable() {
			return "null.Int"
		}
		return "int"
	}

	if strings.HasPrefix(col.Type, "decimal") ||
		strings.HasPrefix(col.Type, "numeric") ||
		strings.HasPrefix(col.Type, "double") ||
		strings.HasPrefix(col.Type, "float") {
		if col.IsNullable() {
			return "null.Float"
		}
		return "float64"
	}

	if strings.HasPrefix(col.Type, "bit") ||
		strings.HasPrefix(col.Type, "tinyint") {
		if col.IsNullable() {
			return "null.Bool"
		}
		return "bool"
	}

	if strings.HasPrefix(col.Type, "timestamp") ||
		strings.HasPrefix(col.Type, "datetime") ||
		strings.HasPrefix(col.Type, "date") {
		if col.IsNullable() {
			return "null.Time"
		}
		return "time.Time"
	}

	return "unknow type"
}
