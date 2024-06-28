package repositories

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConnection(connectString string) *sql.DB {
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil
	}
	if ok := checkDBConnection(db); !ok {
		return nil
	}
	return db
}

func checkDBConnection(db *sql.DB) bool {
	fmt.Println(db.PingContext(context.Background()))
	return db.PingContext(context.Background()) == nil
}
