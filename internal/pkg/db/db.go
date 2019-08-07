package db

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	"wume-composer/internal/pkg/common/config"
)

/********************/
/*      ERRORS      */
/********************/

var (
	AlreadyInitError = errors.New("db already initialized")
	NotInitError     = errors.New("db wasn't initialized")
)

const (
	uniqueErrorCode = "23505"
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
	// " dbname=" + config.Db.DbName +
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

func queryRow(query string, args ...interface{}) (*sql.Row, error) {
	if dbObj == nil {
		return nil, NotInitError
	}

	return dbObj.QueryRow(query, args...), nil
}

func query(query string, args ...interface{}) (*sql.Rows, error) {
	if dbObj == nil {
		return nil, NotInitError
	}

	return dbObj.Query(query, args...)
}

func exec(query string, args ...interface{}) (sql.Result, error) {
	if dbObj == nil {
		var emptyResult sql.Result
		return emptyResult, NotInitError
	}

	return dbObj.Exec(query, args...)
}

func isExists(tableName string, where string, args ...interface{}) (id int64, err error) {
	row, err := findRowBy(tableName, "id", where, args...)
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

func insert(tableName string, cols string, values string, args ...interface{}) (id int64, err error) {
	row, err := queryRow("INSERT INTO "+tableName+" ("+cols+") VALUES ("+values+") RETURNING id", args...)
	if err != nil {
		return
	}
	err = row.Scan(&id)
	if err != nil {
		return
	}
	return id, nil
}

func findRowBy(tableName string, cols string, where string, args ...interface{}) (*sql.Row, error) {
	if where == "" {
		where = "1"
	}
	return queryRow("SELECT "+cols+" FROM "+tableName+" WHERE "+where, args...)
}

func findRowsBy(tableName string, cols string, where string, args ...interface{}) (*sql.Rows, error) {
	if dbObj == nil {
		return nil, NotInitError
	}

	if where == "" {
		where = "1"
	}
	return query("SELECT "+cols+" FROM "+tableName+" WHERE "+where, args...)
}

func updateBy(tableName string, set string, where string, args ...interface{}) (int64, error) {
	if dbObj == nil {
		return 0, NotInitError
	}

	if where == "" {
		where = "1"
	}
	result, err := exec("UPDATE "+tableName+" SET "+set+" WHERE "+where, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func removeBy(tableName string, where string, args ...interface{}) (int64, error) {
	if dbObj == nil {
		return 0, NotInitError
	}

	if where == "" {
		where = "1"
	}
	result, err := exec("DELETE FROM "+tableName+" WHERE "+where, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func truncate(tableName string) error {
	if dbObj == nil {
		return NotInitError
	}

	_, err := exec("TRUNCATE TABLE " + tableName)
	if err != nil {
		return err
	}
	return nil
}
