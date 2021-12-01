package databases

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
)

type DBSchool struct {
	Conn *sql.DB
}

var (
	DBInit     sync.Once
	DBInstance *DBSchool
)

func NewDBSchool() (*DBSchool, error) {
	var (
		conn *sql.DB
		err  error
	)

	DBInit.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			"localhost",
			"5432",
			"postgres",
			"password",
			"db_school",
		)
		if conn, err = SQLConnect("postgres", dsn); err != nil {
			return
		}

		DBInstance = &DBSchool{
			Conn: conn,
		}
	})

	if err != nil {
		return nil, err
	}

	if DBInstance == nil {
		return nil, errors.New("DB problem")
	}

	return DBInstance, nil
}

func SQLConnect(driver, dsn string) (conn *sql.DB, err error) {
	conn, err = sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, err
}
