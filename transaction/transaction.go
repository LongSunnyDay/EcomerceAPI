package transaction

import (
	"database/sql"
	"go-api-ws/helpers"
)

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type TxFn func(Transaction) error

func WithTransaction(db *sql.DB, fn TxFn) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			err := tx.Rollback()
			helpers.CheckErr(err)
			panic(p)
		} else if err != nil {
			err := tx.Rollback()
			helpers.CheckErr(err)
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(tx)
	return err
}
