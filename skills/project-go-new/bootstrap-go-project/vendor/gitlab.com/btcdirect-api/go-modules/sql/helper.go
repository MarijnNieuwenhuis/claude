package sql

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// TODO Move to pkg and add comments
func ExecuteInsert(conn DBConnection, table string, data interface{}) (int64, error) {

	db := conn.DB(true)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfunc()

	query, err := generateInsertQuery(table, data)
	if err != nil {
		return 0, err
	}

	res, err := db.NamedExecContext(ctx, query, data)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func ExecuteUpdate(conn DBConnection, table string, data interface{}) error {

	db := conn.DB(true)

	ctx, cancelfuc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfuc()

	query, err := generateUpdateQuery(table, data)

	if err != nil {
		return err
	}

	if _, err := db.NamedExecContext(ctx, query, data); err != nil {
		return err
	}

	return nil
}

func ExecuteGet(conn DBConnection, table string, id int64, data interface{}) (interface{}, error) {

	db := conn.DB(true)

	ctx, cancelfuc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfuc()
	row, err := db.NamedQueryContext(ctx, fmt.Sprintf("SELECT * FROM %v WHERE id = :id", table), map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	row.Next()

	if err = row.StructScan(data); err != nil {
		return nil, err
	}

	return data, nil
}

func generateInsertQuery(tableName string, data interface{}) (string, error) {
	value := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	// If the value is a pointer, dereference it to get the actual struct value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		typ = typ.Elem()
	}

	if value.Kind() != reflect.Struct {
		return "", fmt.Errorf("data is not a struct")
	}

	var columns []string
	var placeholders []string

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("db")
		sqlTag := field.Tag.Get("sql")

		if tag == "" || sqlTag == "" {
			continue // Skip fields without db tag or no sql tag
		}

		if sqlTag == "update" {
			continue // Skip fields with sql update tag
		}

		columns = append(columns, tag)
		placeholders = append(placeholders, ":"+tag)
	}

	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s);", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	return query, nil
}

func generateUpdateQuery(tableName string, data interface{}) (string, error) {
	value := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	// If the value is a pointer, dereference it to get the actual struct value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		typ = typ.Elem()
	}

	if value.Kind() != reflect.Struct {
		return "", fmt.Errorf("data is not a struct")
	}

	var columns []string

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("db")
		sqlTag := field.Tag.Get("sql")

		if tag == "" || sqlTag == "" {
			continue // Skip fields without db tag
		}

		if sqlTag == "insert" {
			continue // Skip fields with sql insert tag
		}

		value := value.Field(i).Interface()
		if value != reflect.Zero(field.Type).Interface() {
			columns = append(columns, fmt.Sprintf("%s=:%s", tag, tag))
		}
	}

	if len(columns) == 0 {
		return "", fmt.Errorf("no columns to update")
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %s;", tableName, strings.Join(columns, ", "), ":"+typ.Field(0).Tag.Get("db"))

	return query, nil
}
