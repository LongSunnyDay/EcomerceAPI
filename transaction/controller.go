package transaction

import (
	"database/sql"
	"go-api-ws/helpers"
)

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
