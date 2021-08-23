package services

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/silverlovesl/orm-entity-gen/models"
	"github.com/silverlovesl/orm-entity-gen/utils"
)

type GormService struct {
}

func NewGormService() *GormService {
	return &GormService{}
}

func (serv *GormService) GenerateEntity(tableName, packageName string, cols []models.ColumnDesc) []string {
	result := []string{}
	result = append(result, fmt.Sprintf("package %s\n", packageName))
	result = append(result, fmt.Sprintf("type %s struct {\n", tableName))
	// result = append(result, "gorm.Model\n")

	for _, col := range cols {
		cField := strcase.ToCamel(strings.ToLower(col.Field))
		cField = utils.ConvertSpecifyWordToFullUpperCase(cField, "Id", "ID")
		cField = utils.ConvertSpecifyWordToFullUpperCase(cField, "Url", "URL")
		cType := col.Get2GoType()
		primaryKey := ""
		if col.IsPK() {
			primaryKey = "primaryKey; "
		}
		isNull := "not null"
		if col.IsNullable() {
			isNull = "null"
		}

		result = append(result, fmt.Sprintf("%s %s `gorm:\"%scolumn:%s; type:%s %s;\"`\n", cField, cType, primaryKey, col.Field, col.Type, isNull))
	}

	result = append(result, "}")

	return result
}
