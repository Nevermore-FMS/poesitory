package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"
)

var db *sql.DB
var node *snowflake.Node

func CloseDb() {
	db.Close()
}

func Init() error {
	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL
	connStr := os.Getenv("DEV_DB_URI")
	if len(connStr) == 0 {
		connStr = fmt.Sprintf("postgres://postgres:%s@db:5432/poesitory?sslmode=disable", os.Getenv("POESITORY_SECRET"))
	}
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	db = dbConn
	db.SetMaxOpenConns(50)

	snowflake.Epoch = 1628042817000
	node, err = snowflake.NewNode(1)
	return err
}
