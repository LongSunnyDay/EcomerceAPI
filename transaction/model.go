package transaction

import "database/sql"

type (
	TxFn func(Transaction) error

	Transaction interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		Prepare(query string) (*sql.Stmt, error)
		Query(query string, args ...interface{}) (*sql.Rows, error)
		QueryRow(query string, args ...interface{}) *sql.Row
	}

	PipelineStmt struct {
		query string
		args  []interface{}
	}
)
