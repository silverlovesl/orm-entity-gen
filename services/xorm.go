package services

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/silverlovesl/orm-entity-gen/models"
	"github.com/silverlovesl/orm-entity-gen/utils"
)

type XormService struct {
}

func NewXormService() *XormService {
	return &XormService{}
}

func (serv *XormService) GenerateEntity(tableName, packageName string, cols []models.ColumnDesc) []string {
	result := []string{}
	result = append(result, fmt.Sprintf("package %s\n", packageName))
	result = append(result, fmt.Sprintf("type %s struct {\n", tableName))

	for _, col := range cols {
		cField := strcase.ToCamel(strings.ToLower(col.Field))
		cField = utils.ConvertSpecifyWordToFullUpperCase(cField, "Id", "ID")
		cField = utils.ConvertSpecifyWordToFullUpperCase(cField, "Url", "URL")
		cType := col.Get2GoType()
		primaryKey := ""
		if col.IsPK() {
			primaryKey = "pk "
		}

		result = append(result, fmt.Sprintf("%s %s `xorm:\"%s%s '%s'\"`\n", cField, cType, primaryKey, col.Type, col.Field))
	}

	result = append(result, "}")

	return result
}
