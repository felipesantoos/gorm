package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"runtime/debug"
)

type Driver int

const (
	Sqlite = iota + 1
	Mysql
	Postgres
)

type DB struct {
	conn   *sqlx.DB
	lock   Locker
	driver Driver
}

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

type (
	Queryer interface {
		Query(query string, args ...any) (*sql.Rows, error)
		QueryRow(query string, args ...any) *sql.Row
	}
	Binder interface {
		BindNamed(query string, arg any) (string, []interface{}, error)
	}
	Execer interface {
		Queryer
		Exec(query string, args ...any) (sql.Result, error)
	}
)

func (db *DB) View(fn func(Queryer, Binder) error) error {
	db.lock.RLock()
	err := fn(db.conn, db.conn)
	db.lock.RUnlock()
	return err
}

func (db *DB) Lock(fn func(Execer, Binder) error) error {
	db.lock.Lock()
	err := fn(db.conn, db.conn)
	db.lock.Unlock()
	return err
}

func (db *DB) Update(fn func(Execer, Binder) error) (err error) {
	db.lock.Lock()
	defer db.lock.Unlock()
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			debug.PrintStack()
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(tx, db.conn)
	return err
}
