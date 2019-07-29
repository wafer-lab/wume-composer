package db

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var (
	NotInit     = errors.New("db wasn't initialized")
	AlreadyInit = errors.New("db already initialized")
)

var dbObj *sql.DB

func Ping() error {
	return dbObj.Ping()
}

func Open() (err error) {
	if dbObj != nil {
		return AlreadyInit
	}
	dbObj, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/")
	if err != nil {
		return
	}
	err = dbObj.Ping()
	return
}

func Close() error {
	return dbObj.Close()
}

func QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	if dbObj == nil {
		return nil, NotInit
	}

	return dbObj.QueryRow(query, args...), nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	if dbObj == nil {
		return nil, NotInit
	}

	return dbObj.Query(query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	if dbObj == nil {
		var emptyResult sql.Result
		return emptyResult, NotInit
	}

	return dbObj.Exec(query, args...)
}

func isExists(dbName string, tableName string, where string, args ...interface{}) (id int64, err error) {
	row, err := findRowBy(dbName, tableName, "id", where, args...)
	if err != nil {
		return
	}
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return
	}
	return id, nil
}

func insert(dbName string, tableName string, cols string, values string, args ...interface{}) (int64, error) {
	result, err := Exec("INSERT INTO "+dbName+"."+tableName+" ("+cols+") VALUES ("+values+")", args...)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func findRowBy(dbName string, tableName string, cols string, where string, args ...interface{}) (*sql.Row, error) {
	if where == "" {
		where = "1"
	}
	return QueryRow("SELECT "+cols+" FROM "+dbName+"."+tableName+" WHERE "+where, args...)
}

// For future use
//
// func findRowsBy(dbName string, tableName string, cols string, where string, args ...interface{}) (*sql.Rows, error) {
// 	if dbObj == nil {
// 		return nil, NotInit
// 	}
//
// 	if where == "" {
// 		where = "1"
// 	}
// 	return Query("SELECT "+cols+" FROM "+dbName+"."+tableName+" WHERE "+where, args...)
// }

func updateBy(dbName string, tableName string, set string, where string, args ...interface{}) (int64, error) {
	if dbObj == nil {
		return 0, NotInit
	}

	if where == "" {
		where = "1"
	}
	result, err := Exec("UPDATE "+dbName+"."+tableName+" SET "+set+" WHERE "+where, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func removeBy(dbName string, tableName string, where string, args ...interface{}) (int64, error) {
	if dbObj == nil {
		return 0, NotInit
	}

	if where == "" {
		where = "1"
	}
	result, err := Exec("DELETE FROM "+dbName+"."+tableName+" WHERE "+where, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func truncate(dbName string, tableName string) error {
	if dbObj == nil {
		return NotInit
	}

	_, err := Exec("TRUNCATE TABLE " + dbName + "." + tableName)
	if err != nil {
		return err
	}
	return nil
}
