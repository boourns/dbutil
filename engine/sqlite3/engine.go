package sqlite3

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Engine struct {
	*sql.DB
}

func (e Engine) TableExists(name string) bool {
	// sqlite3 syntax
	var table string
	err := e.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name=?;`, name).Scan(&table)
	if table == name && err == nil {
		return true
	}

	return false
}

// ListTables returns a list of tables
func (e Engine) Tables() ([]string, error) {
	rows, err := e.Query("SELECT name FROM sqlite_master WHERE type='table';")
	result := []string{}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		result = append(result, tableName)
	}

	return result, nil
}

func sqlType(f reflect.StructField) string {
	var t string
	switch f.Type.Name() {
	case "string":
		t = "VARCHAR(255)"
	case "int", "int64":
		t = "INTEGER"
	default:
		log.Fatalf("Unknown SQL type for go field %s", f.Type.Name())
	}

	if f.Name == "ID" || f.Name == "id" || f.Name == "Id" {
		t = fmt.Sprintf("%s PRIMARY KEY AUTOINCREMENT", t)
	}
	return t
}

func (e Engine) CreateTable(i interface{}) error {
	t := reflect.TypeOf(i)

	if t.Kind() != reflect.Struct {
		return errors.New("Parameter to CreateTable must be a struct")
	}

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields = append(fields, fmt.Sprintf("%s %s", field.Name, sqlType(field)))
	}
	sql := fmt.Sprintf("CREATE TABLE %s ( %s )", t.Name(), strings.Join(fields, ","))

	_, err := e.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
