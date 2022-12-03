package godatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/db_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
