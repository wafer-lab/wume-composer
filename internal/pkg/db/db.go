package db

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"wume-composer/internal/pkg/config"
)

/********************/
/*      ERRORS      */
/********************/

var (
	AlreadyInitError = errors.New("db already initialized")
	NotInitError     = errors.New("db wasn't initialized")
)

/********************/
/*  BASE FUNCTIONS  */
/********************/

var dbObj *sql.DB

func Ping() error {
	return dbObj.Ping()
}

func Open() (err error) {
	if dbObj != nil {
		return AlreadyInitError
	}

	source := "host=" + config.Db.Host +
		" port=" + config.Db.Port +
		" dbname=" + config.Db.DbName +
		" user=" + config.Db.Username +
		" password=" + config.Db.Password +
		" sslmode=disable"
	dbObj, err = sql.Open("postgres", source)
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
		return nil, NotInitError
	}

	return dbObj.QueryRow(query, args...), nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	if dbObj == nil {
		return nil, NotInitError
	}

	return dbObj.Query(query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	if dbObj == nil {
		var emptyResult sql.Result
		return emptyResult, NotInitError
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

func findRowsBy(dbName string, tableName string, cols string, where string, args ...interface{}) (*sql.Rows, error) {
	if dbObj == nil {
		return nil, NotInitError
	}

	if where == "" {
		where = "1"
	}
	return Query("SELECT "+cols+" FROM "+dbName+"."+tableName+" WHERE "+where, args...)
}

func updateBy(dbName string, tableName string, set string, where string, args ...interface{}) (int64, error) {
	if dbObj == nil {
		return 0, NotInitError
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
		return 0, NotInitError
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
		return NotInitError
	}

	_, err := Exec("TRUNCATE TABLE " + dbName + "." + tableName)
	if err != nil {
		return err
	}
	return nil
}
