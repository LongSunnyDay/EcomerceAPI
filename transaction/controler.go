package transaction

import "database/sql"

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type PipelineStmt struct {
	query string
	args  []interface{}
}

type TxFn func(Transaction) error

func WithTransaction(db *sql.DB, fn TxFn) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = fn(tx)
	return err
}
func NewPipelineStmt(query string, args ...interface{}) *PipelineStmt {
	return &PipelineStmt{query, args}
}

func (ps *PipelineStmt) Exec(tx Transaction) (sql.Result, error) {
	return tx.Exec(ps.query, ps.args)
}

func RunPipeline(tx Transaction, stmts ...*PipelineStmt) (sql.Result, error) {
	var res sql.Result
	var err error

	for _, ps := range stmts {
		res, err = ps.Exec(tx)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
