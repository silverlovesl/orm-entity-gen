package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"database/sql"

	"github.com/gertd/go-pluralize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
	"github.com/silverlovesl/orm-entity-gen/models"
	"github.com/silverlovesl/orm-entity-gen/services"
)

func main() {
	var (
		ormName     string
		packageName string
		output      string
	)
	dbConn := models.DBConnection{}
	flag.StringVar(&dbConn.Username, "user", "", "-user=root")
	flag.StringVar(&dbConn.Password, "password", "", "-password=!QAZ1qaz")
	flag.IntVar(&dbConn.Port, "port", 3306, "-port=3306")
	flag.StringVar(&dbConn.Host, "host", "127.0.0.1", "-host=127.0.0.1")
	flag.StringVar(&dbConn.Database, "database", "", "-database=dbName")
	flag.StringVar(&dbConn.TableName, "table_name", "", "-table_name=tableName")
	flag.StringVar(&packageName, "package_name", "entities", "-package=entities")
	flag.StringVar(&ormName, "orm_name", "xorm", "-orm=xorm")
	flag.StringVar(&output, "output", "xorm", "-output=/Path/your_project/entities")
	flag.Parse()

	db, err := sql.Open("mysql", dbConn.GetMySQLConnString())
	if err != nil {
		log.Fatalf("Can't open database %s\n", err.Error())
	}
	defer db.Close()

	tableService := services.NewTableService(db)
	gormService := services.NewGormService()
	xormService := services.NewXormService()
	pluralize := pluralize.NewClient()

	tableNames, err := tableService.ListAllTables()
	if err != nil {
		log.Fatalf("Failed to get table names %s\n", err.Error())
	}

	var goStructure []string
	tableNameDef := []string{
		fmt.Sprintf("package %s", packageName),
		"\n\n",
		"const (",
	}
	for _, tableName := range tableNames {
		if dbConn.TableName != "all" && dbConn.TableName != tableName {
			continue
		}

		// All 頭文字 COMPANIES = companies
		lTableName := strings.ToLower(tableName)
		// 頭文字小文字 companies = Companies
		lcTableName := strcase.ToCamel(lTableName)
		// 単数にする companies = company
		sTableName := pluralize.Singular(lTableName)
		// 単数 かつ 頭文字大文字 Company
		ucTableName := strcase.ToCamel(sTableName)

		fmt.Printf("Table: %s processing", lcTableName)

		columnDescs, err := tableService.GetColumnDescription(tableName)
		if err != nil {
			log.Fatalf("Failed to get column description %s\n", err.Error())
		}

		if ormName == "xorm" {
			goStructure = xormService.GenerateEntity(ucTableName, packageName, columnDescs)
		} else if ormName == "gorm" {
			goStructure = gormService.GenerateEntity(ucTableName, packageName, columnDescs)
		} else {
			log.Fatalf("Not exists orm name: %s", ormName)
			panic(err)
		}
		filename := fmt.Sprintf("%s/%s.go", output, sTableName)
		writeToFile(filename, goStructure)
		fmt.Println(" | done")

		tableNameDef = append(tableNameDef, fmt.Sprintf("TableName%s = \"%s\"\n", lcTableName, tableName))
	}
	tableNameDef = append(tableNameDef, ")")
	writeToFile(fmt.Sprintf("%s/tables.go", output), tableNameDef)
}

func writeToFile(filename string, text []string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Failed to create file [%s]", filename)
		return err
	}
	defer file.Close()

	for _, line := range text {
		b := []byte(line)
		_, err := file.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
