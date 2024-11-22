package helper

import (
	"database/sql"
	"online-store-golang/errs"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		errs.PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		errs.PanicIfError(errorCommit)
	}
}
