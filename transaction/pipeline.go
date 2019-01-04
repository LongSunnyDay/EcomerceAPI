package transaction

import "database/sql"

type PipelineStmt struct {
	query string
	args  []interface{}
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
