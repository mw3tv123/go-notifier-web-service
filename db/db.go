package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	// MySQL lib
	_ "github.com/go-sql-driver/mysql"
	"github.com/mw3tv123/go-notify/config"
)

var mysqlDB *MySQL

type (
	// MySQL ...
	MySQL struct {
		*sql.DB
	}
	// resultHandler used to handle result from query statement.
	resultHandler func(*sql.Rows) interface{}
)

// createDBConnection Create new DB Context, return error if fail to connect DB.
func createDBConnection(driver, dataSource string) error {
	// Open connection.
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return err
	}
	// Verify connection.
	if err = db.Ping(); err != nil {
		return err
	}
	// Default Setting
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)
	mysqlDB = &MySQL{db}

	return nil
}

// GetDB ...
func GetDB() *MySQL {
	return mysqlDB
}

// Init ConnectToDataBase Connect to main DB (MySQL)
func Init() {
	c := config.GetConfig()

	datasource := fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true", c.GetString("database.username"), c.GetString("database.password"), c.GetString("database.protocol"), c.GetString("database.address"), c.GetString("database.name"))
	err := createDBConnection(c.GetString("database.driver"), datasource)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DONE!")
}

// QueryCommand A wrapper for handle execute query with param and return result from query.
func (db MySQL) QueryCommand(handler resultHandler, queryCmd string, params ...interface{}) (interface{}, error) {
	// For logging purpose, uncomment the follow line
	// log.Print (fmt.Sprintf ("QUERY - SQL - [ %s ] | Params - [ %v ]", queryCmd, params))

	// Prepare statement for query...
	query, err := db.Prepare(queryCmd)
	if err != nil {
		return nil, err
	}
	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	var rows *sql.Rows
	// Check params empty or not.
	// NOTE: If input is null, params still create and hold an element with value of <nil>.
	// So we need to check the type of first element.
	if len(params) > 0 && params[0] != nil {
		rows, err = query.Query(params[:]...)
	} else {
		rows, err = query.Query()
	}
	if err != nil {
		return nil, err
	}

	// Transfer query result to the handler
	result := handler(rows)
	_ = rows.Close()

	return result, nil
}

// ExecuteCommand Execute query, return number rows of effected...
func (db MySQL) ExecuteCommand(executeCmd string, params ...interface{}) (int64, error) {
	// For logging purpose, uncomment the follow line
	// log.Print (fmt.Sprintf ("EXECUTE - SQL - [ %s ] | Params - [ %v ]", executeCmd, params))

	// Prepare statement for query...
	query, err := db.Prepare(executeCmd)
	if err != nil {
		return 0, err
	}
	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	var result sql.Result
	// Check params empty or not.
	// NOTE: If input is null, params still create and hold an element with value of <nil>.
	// So we need to check the type of first element.
	if len(params) > 0 && params[0] != nil {
		result, err = query.Exec(params[:]...)
	} else {
		result, err = query.Exec()
	}
	if err != nil {
		return 0, err
	}
	// Return number of rows effect, the rest depend on caller
	rowEffect, err := result.RowsAffected()

	return rowEffect, err
}
