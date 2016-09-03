package daoutil

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var stmt *sql.Stmt

func init() {
	var err error
	db, err = sql.Open("mysql", "wdeqin:wdeqin@/devdb?charset=utf8")
	if err != nil {
		panic(err)
	}

	stmt, err = db.Prepare("update dev_seq_t set nxtval = last_insert_id(nxtval + 1) where seqnam = ?")
	if err != nil {
		panic(err)
	}
}

func NxtValTrans(seqNam string) int {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec("update dev_seq_t set nxtval = last_insert_id(nxtval + 1) where seqnam = ?", seqNam)
	if err != nil {
		panic(err)
	}

	row := tx.QueryRow("select last_insert_id()")
	if row == nil {
		panic("query sequence failed")
	}

	var nxtVal int = 0
	row.Scan(&nxtVal)

	tx.Commit()

	return nxtVal
}

func NxtVal(seqNam string) int {
	res, err := stmt.Exec(seqNam)
	if err != nil {
		panic(err)
	}

	seq, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(seq)

}
